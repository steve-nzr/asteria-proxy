package consumer

import (
	"encoding/json"
	"os"

	"github.com/steve-nzr/asteria-proxy/pkg/logger"
	"github.com/steve-nzr/asteria-proxy/pkg/rabbitmq/channel"
	"github.com/steve-nzr/asteria-proxy/pkg/server"
	"github.com/streadway/amqp"
)

// GetConsumeChannel gets the go's channel which relays rabbitmq events from a specific exchange
func GetConsumeChannel(exchange channel.Exchange) <-chan amqp.Delivery {
	ch, err := channel.Channel.Consume(getQueue(exchange).Name, "", false, false, false, false, nil)
	exitOnError(err)
	return ch
}

func getQueue(exchange channel.Exchange) *amqp.Queue {
	q, err := channel.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	exitOnError(err)

	err = channel.Channel.QueueBind(
		q.Name,
		"",
		exchange.ToString(),
		false,
		nil,
	)
	exitOnError(err)
	logger.Debug("Queue %s binded to the channel", exchange.ToString())
	return &q
}

func exitOnError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func consumeDisconnect(s *server.Server, ch <-chan amqp.Delivery) {
	logger.Debug("Consuming disconnect queue")
	for m := range ch {
		var msg SendDisconnect
		if err := json.Unmarshal(m.Body, &msg); err != nil {
			logger.Error(err.Error())
			m.Nack(false, false)
			continue
		}

		logger.Debug("Disconnecting client %s", msg.ClientID)
		if err := s.DisconnectClient(msg.ClientID); err != nil {
			m.Nack(false, false)
		}

		m.Ack(false)
	}
}

func consumeMessage(s *server.Server, ch <-chan amqp.Delivery) {
	logger.Debug("Consuming messages queue")
	for m := range ch {
		var msg SendMessage
		if err := json.Unmarshal(m.Body, &msg); err != nil {
			logger.Error(err.Error())
			m.Nack(false, false)
			continue
		}

		logger.Debug("Sending a message of length %d to %d clients", len(msg.Data), len(msg.Recipients))
		s.SendMessage(msg.Recipients, msg.Data)

		m.Ack(false)
	}
}

// Consume back-end queues (Disconnect & Messages)
func Consume(s *server.Server) {
	go consumeDisconnect(s, GetConsumeChannel(channel.ClientDisconnectOut))
	go consumeMessage(s, GetConsumeChannel(channel.ClientMessageOut))
}

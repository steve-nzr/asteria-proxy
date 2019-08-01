package consumer

import (
	"encoding/json"
	"os"

	"github.com/steve-nzr/asteria-proxy/internal/channel"
	"github.com/steve-nzr/asteria-proxy/pkg/logger"
	"github.com/steve-nzr/asteria-proxy/pkg/server"
	"github.com/streadway/amqp"
)

var queues = map[channel.Exchange]*amqp.Queue{
	channel.ClientDisconnectOut: getQueue(channel.ClientDisconnectOut),
	channel.ClientMessageOut:    getQueue(channel.ClientMessageOut),
}

func exitOnError(err error) {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
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

func getConsumeChannel(exchange channel.Exchange) <-chan amqp.Delivery {
	ch, err := channel.Channel.Consume(queues[exchange].Name, "", false, false, false, false, nil)
	exitOnError(err)
	return ch
}

func consumeDisconnect(s *server.Server, ch <-chan amqp.Delivery) {
	logger.Debug("Consuming disconnect queue")
	for m := range ch {
		var msg OnDisconnect
		if err := json.Unmarshal(m.Body, &msg); err != nil {
			logger.Error(err.Error())
			m.Nack(false, false)
			continue
		}

		m.Ack(false)
	}
}

func consumeMessage(s *server.Server, ch <-chan amqp.Delivery) {
	logger.Debug("Consuming messages queue")
	for m := range ch {
		var msg OnMessage
		if err := json.Unmarshal(m.Body, &msg); err != nil {
			logger.Error(err.Error())
			m.Nack(false, false)
			continue
		}

		m.Ack(false)
	}
}

// Consume back-end queues (Disconnect & Messages)
func Consume(s *server.Server) {
	go consumeDisconnect(s, getConsumeChannel(channel.ClientDisconnectOut))
	go consumeMessage(s, getConsumeChannel(channel.ClientMessageOut))
}

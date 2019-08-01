package publisher

import (
	"encoding/json"

	"github.com/steve-nzr/asteria-proxy/internal/channel"
	"github.com/streadway/amqp"
)

// Publish a message to the given exchange queues
func Publish(exchange channel.Exchange, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return channel.Channel.Publish(
		exchange.ToString(), // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

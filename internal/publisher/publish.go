package publisher

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// Publish a message to the given exchange queues
func Publish(exchange Exchange, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return channel.Publish(
		(string)(exchange), // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

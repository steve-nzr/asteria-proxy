package channel

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/steve-nzr/asteria-proxy/pkg/rabbitmq"
)

// Channel rabbitmq that contains topics
var Channel = initialize()

func initialize() *amqp.Channel {
	ch, err := rabbitmq.Connection.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}
	//defer ch.Close()

	for _, exchange := range exchanges {
		err = ch.ExchangeDeclare(
			(string)(exchange), // name
			"fanout",           // type
			true,               // durable
			false,              // auto-deleted
			false,              // internal
			false,              // no-wait
			nil,                // arguments
		)
	}

	return ch
}

package publisher

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/steve-nzr/asteria-proxy/pkg/rabbitmq"
)

// Exchange rabbitmq
type Exchange string

const (
	// ClientConnected exchange
	ClientConnected Exchange = "client_connected"
	// ClientDisonnected exchange
	ClientDisonnected Exchange = "client_disconnected"
	// ClientMessage exchange
	ClientMessage Exchange = "client_message"
)

var exchanges = []Exchange{
	ClientConnected,
	ClientDisonnected,
	ClientMessage,
}

var channel = initialize()

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

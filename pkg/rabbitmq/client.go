package rabbitmq

import (
	"os"

	"github.com/streadway/amqp"
)

// Connection holds an active connection to rabbitmq
var Connection = initialize()

func initialize() *amqp.Connection {
	conn, err := amqp.Dial(os.Getenv("AMQP_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	return conn
}

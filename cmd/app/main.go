package main

import (
	"github.com/steve-nzr/asteria-proxy/pkg/rabbitmq/consumer"
	"github.com/steve-nzr/asteria-proxy/pkg/server"
)

func main() {
	server := new(server.Server)
	go consumer.Consume(server)
	server.Start()
}

package channel

import (
	"fmt"
	"os"

	"github.com/steve-nzr/asteria-proxy/pkg/logger"
)

// Exchange rabbitmq
type Exchange string

// ToString returns Exchange as string type
func (e Exchange) ToString() string {
	return (string)(e)
}

var (
	// ClientConnected exchange
	ClientConnected = getExchangeFromEnv("AMQP_TOPIC_IN_CONNECTED")
	// ClientDisonnected exchange
	ClientDisonnected = getExchangeFromEnv("AMQP_TOPIC_IN_DISCONNECTED")
	// ClientMessage exchange
	ClientMessage = getExchangeFromEnv("AMQP_TOPIC_IN_MESSAGE")
	// ClientDisconnectOut exchange
	ClientDisconnectOut = getExchangeFromEnv("AMQP_TOPIC_OUT_DISCONNECT")
	// ClientMessageOut exchange
	ClientMessageOut = getExchangeFromEnv("AMQP_TOPIC_OUT_MESSAGE")
)

// exchanges is list of topics
var exchanges = []Exchange{
	ClientConnected,
	ClientDisonnected,
	ClientMessage,
	ClientDisconnectOut,
	ClientMessageOut,
}

func getExchangeFromEnv(key string) Exchange {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		logger.Error(fmt.Errorf("cannot find environment variable %s, or its empty", key).Error())
		os.Exit(1)
	}
	return (Exchange)(v)
}

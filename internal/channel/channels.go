package channel

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

// exchanges is list of topics
var exchanges = []Exchange{
	ClientConnected,
	ClientDisonnected,
	ClientMessage,
}

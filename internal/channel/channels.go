package channel

// Exchange rabbitmq
type Exchange string

// ToString returns Exchange as string type
func (e Exchange) ToString() string {
	return (string)(e)
}

const (
	// ClientConnected exchange
	ClientConnected Exchange = "in_client_connected"
	// ClientDisonnected exchange
	ClientDisonnected Exchange = "in_client_disconnected"
	// ClientMessage exchange
	ClientMessage Exchange = "in_client_message"
	// ClientDisconnectOut exchange
	ClientDisconnectOut Exchange = "out_client_disconnect"
	// ClientMessageOut exchange
	ClientMessageOut Exchange = "out_client_message"
)

// exchanges is list of topics
var exchanges = []Exchange{
	ClientConnected,
	ClientDisonnected,
	ClientMessage,
	ClientDisconnectOut,
	ClientMessageOut,
}

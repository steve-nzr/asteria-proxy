package publisher

// OnMessage structure, when a new messages arrives
type OnMessage struct {
	ClientID       string `json:"client_id"`
	ReceivedAt     int64  `json:"received_at"`
	LastReceivedAt int64  `json:"last_received_at"`
	Size           int    `json:"size"`
	Data           []byte `json:"data"`
}

// OnConnected structure, when a client just arrived
type OnConnected struct {
	ClientID   string `json:"client_id"`
	ReceivedAt int64  `json:"received_at"`
}

// OnDisconnected structure, when a client disconnected or connection has been lost to it
type OnDisconnected struct {
	ClientID   string `json:"client_id"`
	ReceivedAt int64  `json:"received_at"`
}

package consumer

// OnDisconnect structure, when back-end notifies us to disconnect a client
type OnDisconnect struct {
	ClientID string `json:"client_id"`
	Reason   string `json:"reason,omitempty"`
}

// OnMessage structure, when back-end notifies us to send a message to one, multiple or all clients
// empty recipients means broadcast
type OnMessage struct {
	Data       []byte   `json:"data"`
	Recipients []string `json:"recipients,omitempty"`
}

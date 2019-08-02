package consumer

// SendDisconnect structure, when back-end notifies us to disconnect a client
type SendDisconnect struct {
	ClientID string `json:"client_id"`
	Reason   string `json:"reason,omitempty"`
}

// SendMessage structure, when back-end notifies us to send a message to one, multiple or all clients
// empty recipients means broadcast
type SendMessage struct {
	Data       []byte   `json:"data"`
	Recipients []string `json:"recipients,omitempty"`
}

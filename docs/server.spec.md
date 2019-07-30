# Server SPECS

Client : {
    id string
}

OnConnect -> {
    client_id_: string
    received_at_: int64
}

OnDisconnect -> {
    client_id_: string
    received_at_: int64
}

OnMessage -> {
    client_id_: string
    received_at_: int64
    last_received_at_: int64
    size: int32
    data: []byte
}

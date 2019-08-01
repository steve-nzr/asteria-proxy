# Server SPECS

## Global informations
- Time variables are in Unix timestamp (time elaspsed since 1 January 1970) in nanoseconds.
- ClientID are V4 UUID generated at connection.

## Messages structures
Client structure, used only by the proxy but we found useful to show you what is was like.
```go
{
    id string
}
```

---

**Connection** message format, sent when a new client connects successfully to the proxy.
- Topic : *in_client_connected*
```go
{
    client_id: string
    received_at: int64
}
```

---

**Disconnection** message format, sent when a client disconnects from the proxy, intentionally or not (connection lost for example).
- Topic : *in_client_disconnected*
```go
{
    client_id: string
    received_at: int64
}
```

---

**Message** format, sent when a client emits a new message.

*last_received_at* is the last message's time.
- Topic : *in_client_message*
```go
{
    client_id: string
    received_at: int64
    last_received_at: int64
    size: int32
    data: []byte
}
```

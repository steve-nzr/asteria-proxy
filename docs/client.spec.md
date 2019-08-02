 # Client SPECS

## Global informations
- ClientID are V4 UUID generated at connection.

## Messages structures
**Disconnect** message format, send it when you need to disconnect a specific client.
- Topic : *out_client_disconnect*
```go
{
    client_id: string
    reason: string
}
```

---

**Message** format, send it when you need to send a message to one, multiple or all clients.
To broadcast a message, leave recipients empty.
- Topic : *out_client_message*
```go
{
    data: []byte
    recipients: []string
}
```

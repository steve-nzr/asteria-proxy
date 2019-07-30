package server

import (
	"net"
	"time"

	"github.com/steve-nzr/asteria-proxy/internal/publisher"
	"github.com/steve-nzr/asteria-proxy/pkg/logger"
)

// Client holds a socket + an UUID as ClientID
type Client struct {
	Conn     net.Conn
	ClientID string
}

func (c *Client) sendMessage(msg *publisher.OnMessage) {
	publisher.Publish(publisher.ClientMessage, msg)
	logger.Debug("Message with length of %d arrived from client %s !", msg.Size, c.ClientID)
}

func (c *Client) sendDisconnected(msg *publisher.OnDisconnected) {
	publisher.Publish(publisher.ClientMessage, msg)
	logger.Debug("Client %s disconnected !", c.ClientID)
}

func (c *Client) handle() {
	defer c.Conn.Close()

	publisher.Publish(publisher.ClientMessage, &publisher.OnConnected{
		ClientID:   c.ClientID,
		ReceivedAt: timeNow(),
	})

	var last int64

	logger.Debug("Client %s connected !", c.ClientID)
	for {
		c.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		buf := make([]byte, bufferSize)
		len, err := c.Conn.Read(buf)
		now := timeNow()

		if err != nil {
			netErr, ok := err.(net.Error)
			if !ok || !netErr.Timeout() {
				c.sendDisconnected(&publisher.OnDisconnected{
					ClientID:   c.ClientID,
					ReceivedAt: now,
				})
				break
			}
			continue
		}

		c.sendMessage(&publisher.OnMessage{
			ClientID:       c.ClientID,
			Data:           buf,
			ReceivedAt:     now,
			Size:           len,
			LastReceivedAt: last,
		})

		last = now
	}
}

func timeNow() int64 {
	return time.Now().Unix()
}

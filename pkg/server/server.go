package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/steve-nzr/asteria-proxy/pkg/logger"
	"github.com/steve-nzr/asteria-proxy/pkg/rabbitmq/publisher"
)

// Server main struct
type Server struct {
	Clients []*Client
	stopped bool
}

// Start the server
func (s *Server) Start() {
	l, err := net.Listen(os.Getenv("SERVER_TYPE"), fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
		return
	}

	logger.Info("Server is started !")
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Print(err.Error())
				continue
			}

			c := new(Client)
			*c = Client{
				Conn:     conn,
				ClientID: uuid.New().String(),
			}
			s.Clients = append(s.Clients, c)
			go c.handle()
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGABRT)
	<-sigs
	s.stop(l)
}

// DisconnectClient by its ID
func (s *Server) DisconnectClient(clientID string) error {
	for _, c := range s.Clients {
		if c.ClientID == clientID {
			c.sendDisconnected(&publisher.OnDisconnected{
				ClientID:   clientID,
				ReceivedAt: timeNow(),
			})
			return c.Conn.Close()
		}
	}

	err := fmt.Errorf("client %s not found", clientID)
	logger.Error(err.Error())
	return err
}

// SendMessage to client(s)
func (s *Server) SendMessage(clients []string, data []byte) {
	broadcast := len(clients) == 0
	for _, c := range s.Clients {
		if broadcast || find(clients, c.ClientID) != -1 {
			if _, err := c.Conn.Write(data); err != nil {
				logger.Warning("cannot write to client %s", c.ClientID)
			}
			logger.Debug("wrote a message to client %s", c.ClientID)
		}
	}
}

func (s *Server) stop(l net.Listener) {
	now := timeNow()
	for _, c := range s.Clients {
		c.sendDisconnected(&publisher.OnDisconnected{
			ClientID:   c.ClientID,
			ReceivedAt: now,
		})
		c.Conn.Close()
	}

	l.Close()
	logger.Info("Server is stopping...")
	os.Exit(0)
}

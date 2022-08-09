package internal

import (
	"net"
	"time"
)

// client is our user application handler.
type client struct {
	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	return &client{
		network: network{
			connection: conn,
		},
	}
}

// Publish will send a message to broker server.
func (c *client) Publish(data []byte) error {
	err := c.network.send(data)
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	return nil
}

func (c *client) Subscribe(handler MessageHandler) {
	go func() {
		var (
			err    error
			buffer = make([]byte, 1024)
		)

		for {
			buffer, err = c.network.get(buffer)
			if err != nil {
				break
			}

			handler(buffer)
		}
	}()
}

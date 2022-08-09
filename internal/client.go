package internal

import (
	"net"
	"time"
)

// client is our user application handler.
type client struct {
	communicateChannel chan []byte
	subscribe          bool

	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	c := &client{
		communicateChannel: make(chan []byte),
		subscribe:          false,

		network: network{
			connection: conn,
		},
	}

	// starting data reader
	go c.readDataFromServer()

	return c
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

// readDataFromServer gets all data from server.
func (c *client) readDataFromServer() {
	var (
		err    error
		buffer = make([]byte, 1024)
	)

	for {
		buffer, err = c.network.get(buffer)
		if err != nil {
			break
		}

		if c.subscribe {
			c.communicateChannel <- buffer
		}
	}
}

// Subscribe subscribes over broker.
func (c *client) Subscribe(handler MessageHandler) {
	go func() {
		c.subscribe = true

		for {
			select {
			case data := <-c.communicateChannel:
				handler(data)
			}
		}
	}()
}

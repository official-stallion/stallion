package internal

import (
	"net"
	"time"
)

// client is our user application handler.
type client struct {
	// communication channel allows a client to make
	// a connection channel between read data and subscribers
	communicateChannel chan []byte

	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	c := &client{
		communicateChannel: make(chan []byte),

		network: network{
			connection: conn,
		},
	}

	// starting data reader
	go c.readDataFromServer()

	return c
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

		m, _ := decodeMessage(buffer)

		switch m.Type {
		case Message:
			c.communicateChannel <- []byte(m.Data)
		}
	}
}

// Publish will send a message to broker server.
func (c *client) Publish(data []byte) error {
	err := c.network.send(encodeMessage(newMessage(Message, data)))
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	return nil
}

// Subscribe subscribes over broker.
func (c *client) Subscribe(handler MessageHandler) {
	go func() {
		_ = c.network.send(encodeMessage(newMessage(Subscribe, nil)))

		time.Sleep(10 * time.Millisecond)

		for {
			select {
			case data := <-c.communicateChannel:
				handler(data)
			}
		}
	}()
}

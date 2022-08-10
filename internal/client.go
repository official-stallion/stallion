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
	// terminate channel is used to close a subscribe channel
	terminateChannel chan int

	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	c := &client{
		communicateChannel: make(chan []byte),
		terminateChannel:   make(chan int),

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
		case Text:
			c.communicateChannel <- m.Data
		}
	}
}

// Publish will send a message to broker server.
func (c *client) Publish(data []byte) error {
	err := c.network.send(encodeMessage(newMessage(Text, data)))
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	return nil
}

// Subscribe subscribes over broker.
func (c *client) Subscribe(handler MessageHandler) {
	_ = c.network.send(encodeMessage(newMessage(Subscribe, []byte(DummyMessage))))
	time.Sleep(10 * time.Millisecond)

	go func() {
		flag := true

		for flag {
			select {
			case data := <-c.communicateChannel:
				handler(data)
			case <-c.terminateChannel:
				flag = false
			}
		}
	}()
}

func (c *client) Unsubscribe() {
	_ = c.network.send(encodeMessage(newMessage(Unsubscribe, nil)))
	time.Sleep(10 * time.Millisecond)

	c.terminateChannel <- 1
}

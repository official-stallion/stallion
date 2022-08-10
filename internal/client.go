package internal

import (
	"log"
	"net"
	"time"
)

// client is our user application handler.
type client struct {
	// map of topics
	topics map[string]MessageHandler
	// communication channel allows a client to make
	// a connection channel between read data and subscribers
	communicateChannel chan Message
	// terminate channel is used to close a subscribe channel
	terminateChannel chan int

	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	c := &client{
		topics: make(map[string]MessageHandler),

		communicateChannel: make(chan Message),
		terminateChannel:   make(chan int),

		network: network{
			connection: conn,
		},
	}

	// starting data reader
	go c.readDataFromServer()
	// start listening on channels
	go c.listen()

	return c
}

// readDataFromServer gets all data from server.
func (c *client) readDataFromServer() {
	var (
		buffer = make([]byte, 2048)
	)

	for {
		tmp, err := c.network.get(buffer)
		if err != nil {
			break
		}

		m, _ := decodeMessage(tmp)
		if m.Type == Text {
			c.communicateChannel <- *m
		}
	}

	c.terminateChannel <- 1
}

// listen method watches channels for input data.
func (c *client) listen() {
	for {
		select {
		case data := <-c.communicateChannel:
			c.handle(data)
		case <-c.terminateChannel:
			c.close()
		}
	}
}

// handle will execute the topic handler method.
func (c *client) handle(m Message) {
	c.topics[m.Topic](m.Data)
}

// close will terminate everything.
func (c *client) close() {
	_ = c.network.connection.Close()

	close(c.communicateChannel)
	close(c.terminateChannel)
}

// Publish will send a message to broker server.
func (c *client) Publish(topic string, data []byte) error {
	err := c.network.send(encodeMessage(newMessage(Text, topic, data)))
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	return nil
}

// Subscribe subscribes over broker.
func (c *client) Subscribe(topic string, handler MessageHandler) {
	err := c.network.send(encodeMessage(newMessage(Subscribe, topic, nil)))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Millisecond)

	// set a handler for given topic
	c.topics[topic] = handler
}

func (c *client) Unsubscribe(topic string) {
	_ = c.network.send(encodeMessage(newMessage(Unsubscribe, topic, nil)))
	time.Sleep(10 * time.Millisecond)

	// remove topic and its handler
	delete(c.topics, topic)
}

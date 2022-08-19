package internal

import (
	"net"
	"time"

	"go.uber.org/zap"
)

// we need safety timeout, to prevent send more than
// one message in http request.
const (
	safetyTimeout = 1
)

// client is our user application handler.
type client struct {
	// map of topics
	topics map[string]MessageHandler

	// communication channel allows a client to make
	// a connection channel between read data and subscribers
	communicateChannel chan message

	// terminate channel is used to close a subscribe channel
	terminateChannel chan int

	// network handles the client socket data transfers
	network network
}

// NewClient creates a new client handler.
func NewClient(conn net.Conn) *client {
	c := &client{
		topics:             make(map[string]MessageHandler),
		communicateChannel: make(chan message),
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
	var buffer = make([]byte, 2048)

	for {
		// read data from network
		tmp, er := c.network.get(buffer)
		if er != nil {
			zap.L().Error("failed read data", zap.Error(er))

			break
		}

		// decode message
		if m, err := decodeMessage(tmp); err == nil {
			if m.Type == Text {
				c.communicateChannel <- *m
			}
		} else {
			zap.L().Error("failed in message parse", zap.Error(err))
		}
	}

	// close
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
func (c *client) handle(m message) {
	if handler, ok := c.topics[m.Topic]; ok {
		handler(m.Data)
	}
}

// close will terminate everything.
func (c *client) close() {
	_ = c.network.connection.Close()
}

// Publish will send a message to broker server.
func (c *client) Publish(topic string, data []byte) error {
	err := c.network.send(encodeMessage(newMessage(Text, topic, data)))
	if err != nil {
		return err
	}

	time.Sleep(safetyTimeout * time.Millisecond)

	return nil
}

// Subscribe subscribes over broker.
func (c *client) Subscribe(topic string, handler MessageHandler) {
	// set a handler for given topic
	c.topics[topic] = handler

	// send an http request to broker server
	err := c.network.send(encodeMessage(newMessage(Subscribe, topic, nil)))
	if err != nil {
		zap.L().Fatal("failed send message to broker server", zap.Error(err))
	}

	time.Sleep(safetyTimeout * time.Millisecond)
}

// Unsubscribe removes client from subscribing over a topic.
func (c *client) Unsubscribe(topic string) {
	_ = c.network.send(encodeMessage(newMessage(Unsubscribe, topic, nil)))
	time.Sleep(safetyTimeout * time.Millisecond)

	// remove topic and its handler
	delete(c.topics, topic)
}

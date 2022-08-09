package internal

import (
	"log"
)

// broker handles the message sending and receiving.
type broker struct {
	channels       []chan []byte
	receiveChannel chan []byte
}

// NewBroker generates a broker.
func newBroker(channel chan []byte) *broker {
	return &broker{
		receiveChannel: channel,
	}
}

// start will start our broker logic.
func (b *broker) start() {
	log.Printf("broker start ...\n")

	for {
		select {
		case data := <-b.receiveChannel:
			b.publish(data)
		}
	}
}

// subscribe will add subscribers to our broker.
func (b *broker) subscribe(channel chan []byte) {
	b.channels = append(b.channels, channel)
}

// publish will send a data over channels
func (b *broker) publish(data []byte) {
	for _, channel := range b.channels {
		channel <- data
	}
}

package internal

import (
	"log"
)

// broker handles the message sending and receiving.
type broker struct {
	channels       []workChan
	receiveChannel chan []byte
	statusChannel  chan int
}

type workChan struct {
	id      int
	channel chan []byte
}

// NewBroker generates a broker.
func newBroker(channel chan []byte, status chan int) *broker {
	return &broker{
		receiveChannel: channel,
		statusChannel:  status,
	}
}

// start will start our broker logic.
func (b *broker) start() {
	log.Printf("broker start ...\n")

	go b.unsubscribe()

	for {
		select {
		case data := <-b.receiveChannel:
			b.publish(data)
		}
	}
}

// subscribe will add subscribers to our broker.
func (b *broker) subscribe(channel chan []byte, id int) {
	w := workChan{
		id:      id,
		channel: channel,
	}

	b.channels = append(b.channels, w)
}

// unsubscribe will remove a channel from broker list.
func (b *broker) unsubscribe() {
	for {
		select {
		case id := <-b.statusChannel:
			for index, value := range b.channels {
				if value.id == id {
					b.channels = append(b.channels[:index], b.channels[index+1:]...)

					break
				}
			}
		}
	}
}

// publish will send a data over channels
func (b *broker) publish(data []byte) {
	for _, w := range b.channels {
		w.channel <- data
	}
}

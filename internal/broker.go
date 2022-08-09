package internal

import (
	"log"
)

// broker handles the message sending and receiving.
type broker struct {
	// list of broker workers
	workers []workChan

	// receiveChannel is a public channel between workers and broker
	receiveChannel chan []byte
	// with statusChannel broker manages the workers status
	statusChannel chan int
}

// newBroker generates a broker.
func newBroker(receive chan []byte, status chan int) *broker {
	return &broker{
		receiveChannel: receive,
		statusChannel:  status,
	}
}

// start will start our broker logic.
func (b *broker) start() {
	log.Printf("broker server start ...\n")

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
	b.workers = append(
		b.workers,
		workChan{
			id:      id,
			channel: channel,
		},
	)
}

// unsubscribe will remove a channel from broker list.
func (b *broker) unsubscribe() {
	for {
		select {
		case id := <-b.statusChannel:
			for index, value := range b.workers {
				if value.id == id {
					b.workers = append(b.workers[:index], b.workers[index+1:]...)

					break
				}
			}
		}
	}
}

// publish will send a data over channels.
func (b *broker) publish(data []byte) {
	for _, w := range b.workers {
		w.channel <- data
	}
}

package internal

import (
	"log"
)

// broker handles the message sending and receiving.
type broker struct {
	// list of broker workers
	workers []WorkChan

	// receiveChannel is a public channel between workers and broker
	receiveChannel chan []byte
	// with statusChannel broker manages the workers status
	statusChannel chan WorkChan
}

// newBroker generates a broker.
func newBroker(receive chan []byte, status chan WorkChan) *broker {
	return &broker{
		receiveChannel: receive,
		statusChannel:  status,
	}
}

// start will start our broker logic.
func (b *broker) start() {
	log.Printf("broker server start ...\n")

	go b.listenToWorkers()

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
		WorkChan{
			id:      id,
			channel: channel,
		},
	)
}

// listenToWorkers will update workers based on status channel.
func (b *broker) listenToWorkers() {
	for {
		select {
		case worker := <-b.statusChannel:
			if worker.status {
				b.subscribe(worker.channel, worker.id)
			} else {
				b.removeDeadWorker(worker.id)
			}
		}
	}
}

// removeDeadWorker will remove a channel from broker list.
func (b *broker) removeDeadWorker(id int) {
	for index, value := range b.workers {
		if value.id == id {
			b.workers = append(b.workers[:index], b.workers[index+1:]...)

			break
		}
	}
}

// publish will send a data over channels.
func (b *broker) publish(data []byte) {
	for _, w := range b.workers {
		w.channel <- data
	}
}

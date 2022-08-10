package internal

import (
	"log"
)

// broker handles the message sending and receiving.
type broker struct {
	// list of broker workers
	workers map[string][]WorkChan

	// receiveChannel is a public channel between workers and broker
	receiveChannel chan Message
	// with statusChannel broker manages the workers status
	statusChannel chan WorkChan
}

// newBroker generates a broker.
func newBroker(receive chan Message, status chan WorkChan) *broker {
	return &broker{
		workers: make(map[string][]WorkChan),

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
func (b *broker) subscribe(topic string, channel chan Message, id int) {
	b.workers[topic] = append(
		b.workers[topic],
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
			switch worker.status {
			case SubStatus:
				b.subscribe(worker.topic, worker.channel, worker.id)
			case UnsubStatus:
				b.unsubscribe(worker.topic, worker.id)
			case TerminateStatus:
				b.removeDeadWorker(worker.id)
			}
		}
	}
}

// unsubscribe removes a worker channel from a topic.
func (b *broker) unsubscribe(topic string, id int) {
	for index, value := range b.workers[topic] {
		if value.id == id {
			b.workers[topic] = append(b.workers[topic][:index], b.workers[topic][index+1:]...)

			break
		}
	}
}

// removeDeadWorker will remove a channel from broker list.
func (b *broker) removeDeadWorker(id int) {
	for key := range b.workers {
		b.unsubscribe(key, id)
	}
}

// publish will send a data over channels.
func (b *broker) publish(data Message) {
	for _, w := range b.workers[data.Topic] {
		w.channel <- data
	}
}

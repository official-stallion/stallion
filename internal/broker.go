package internal

import "fmt"

// broker handles the message sending and receiving.
type broker struct {
	// list of broker workers
	workers map[string][]workerChannel

	// receiveChannel is a public channel between workers and broker
	receiveChannel chan message

	// subscribeChannel is a public channel for subscribing workers over a topic
	subscribeChannel chan subscribeChannel

	// unsubscribeChannel is a public channel for unsubscribing workers from a topic
	unsubscribeChannel chan unsubscribeChannel

	// terminateChannel create a channel for dead workers
	terminateChannel chan int
}

// newBroker generates a broker.
func newBroker(
	receive chan message,
	subscribe chan subscribeChannel,
	unsubscribe chan unsubscribeChannel,
	termination chan int,
) *broker {
	return &broker{
		workers:            make(map[string][]workerChannel),
		receiveChannel:     receive,
		subscribeChannel:   subscribe,
		unsubscribeChannel: unsubscribe,
		terminateChannel:   termination,
	}
}

// start will start our broker logic.
func (b *broker) start() {
	logInfo("broker server start")

	// start a process to listen to our workers
	go b.listenToWorkers()

	// wait for receive channel messages
	for {
		data := <-b.receiveChannel

		go b.publish(data)
	}
}

// listenToWorkers will update workers based on status channel.
func (b *broker) listenToWorkers() {
	for {
		select {
		case packet := <-b.subscribeChannel:
			b.subscribe(packet.topic, packet.channel, packet.id)
		case packet := <-b.unsubscribeChannel:
			b.unsubscribe(packet.topic, packet.id)
		case id := <-b.terminateChannel:
			b.removeDeadWorker(id)
		}
	}
}

// subscribe will add subscribers to our broker.
func (b *broker) subscribe(topic string, channel chan message, id int) {
	b.workers[topic] = append(
		b.workers[topic],
		workerChannel{
			id:      id,
			channel: channel,
		},
	)
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

		logInfo("worker removed", fmt.Sprintf("id=%d", id))
	}
}

// publish will send a data over channels.
func (b *broker) publish(data message) {
	for _, w := range b.workers[data.Topic] {
		w.channel <- data

		logInfo("message published", fmt.Sprintf("size=%d", len(data.Data)))
	}
}

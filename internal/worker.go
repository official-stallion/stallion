package internal

import (
	"log"
	"net"
	"time"
)

// worker handles a single client that wants
// to either subscribe or publish.
type worker struct {
	id      int
	network network

	// send channel is used for getting data from broker (its private between broker and worker)
	sendChannel chan Message
	// receive channel is used for sending data to broker (its public)
	receiveChannel chan Message
	// subscribeChannel is a public channel for subscribing workers over a topic
	subscribeChannel chan SubscribeChannel
	// unsubscribeChannel is a public channel for unsubscribing workers from a topic
	unsubscribeChannel chan UnsubscribeChannel
	// terminateChannel create a channel for dead workers
	terminateChannel chan int
}

// newWorker generates a new worker.
func newWorker(
	id int,
	conn net.Conn,
	sen, rec chan Message,
	sub chan SubscribeChannel,
	unsub chan UnsubscribeChannel,
	ter chan int,
) *worker {
	return &worker{
		id: id,
		network: network{
			connection: conn,
		},

		sendChannel:        sen,
		receiveChannel:     rec,
		subscribeChannel:   sub,
		unsubscribeChannel: unsub,
		terminateChannel:   ter,
	}
}

// start will start our worker.
func (w *worker) start() {
	// closing channel after we are done
	defer close(w.sendChannel)

	// start for input data
	go w.arrival()

	// wait for any interrupt in send channel
	for {
		data := <-w.sendChannel

		w.transfer(data)
	}
}

// transfer will send a data byte through handler.
func (w *worker) transfer(data Message) {
	err := w.network.send(encodeMessage(data))
	if err != nil {
		log.Printf("failed to send: %v\n", err)
	}

	time.Sleep(10 * time.Millisecond)
}

// arrival will check for input data from client.
func (w *worker) arrival() {
	var (
		err    error
		buffer = make([]byte, 1024)
	)

	for {
		buffer, err = w.network.get(buffer)
		if err != nil {
			break
		}

		m, er := decodeMessage(buffer)
		if er != nil {
			log.Printf("parse error: %v", er)

			continue
		}

		switch m.Type {
		case Text:
			// passing data to broker channel
			w.receiveChannel <- *m
		case Subscribe:
			// passing subscribe message
			w.subscribeChannel <- SubscribeChannel{
				id:      w.id,
				topic:   m.Topic,
				channel: w.sendChannel,
			}
		case Unsubscribe:
			// passing unsubscribe message
			w.unsubscribeChannel <- UnsubscribeChannel{
				id:    w.id,
				topic: m.Topic,
			}
		}
	}

	// announcing that the worker is done
	w.terminateChannel <- w.id
}

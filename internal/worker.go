package internal

import (
	"log"
	"net"
)

// worker handles a single client that wants
// to either subscribe or publish.
type worker struct {
	id      int
	handler handler

	// send channel is used for getting data from broker (its private between broker and worker)
	sendChannel chan []byte
	// receive channel is used for sending data to broker (its public)
	receiveChannel chan []byte
}

// newWorker generates a new worker.
// id: for worker id.
// conn: http connection over TCP.
// sen: sending channel.
// rec: receive channel.
func newWorker(id int, conn net.Conn, sen, rec chan []byte) *worker {
	return &worker{
		id: id,
		handler: handler{
			conn: conn,
		},
		sendChannel:    sen,
		receiveChannel: rec,
	}
}

// Start will start our worker.
func (w *worker) start() {
	// close send connection
	defer close(w.receiveChannel)

	// start for input data
	go w.receive()

	// wait for any interrupt in send channel
	select {
	case data := <-w.sendChannel:
		w.send(data)
	}
}

// send will send a data byte through handler.
func (w *worker) send(data []byte) {
	if err := w.handler.write(data); err != nil {
		log.Fatalf("[%d] failed to send: %v\n", w.id, err)
	}
}

// receive will check for input data from client.
func (w *worker) receive() {
	for {
		data, err := w.handler.read()
		if err == nil {
			w.receiveChannel <- data
		}
	}
}

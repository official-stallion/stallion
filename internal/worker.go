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

	sendChannel    chan []byte
	receiveChannel chan []byte
}

// NewWorker generates a new worker.
// id: for worker id.
// conn: http connection over TCP.
// sen: sending channel.
// rec: receive channel.
func NewWorker(id int, conn net.Conn, sen, rec chan []byte) *worker {
	return &worker{
		id: id,
		handler: handler{
			conn: conn,
		},
		receiveChannel: rec,
		sendChannel:    sen,
	}
}

// Start will start our worker.
func (w *worker) Start() {
	// close send connection
	defer close(w.sendChannel)

	// start for input data
	go w.receive()

	// wait for any interrupt in send channel
	select {
	case data := <-w.sendChannel:
		w.send(data)
	}
}

func (w *worker) send(data []byte) {
	if err := w.handler.Write(data); err != nil {
		log.Fatalf("[%d] failed to send: %v\n", w.id, err)
	}
}

func (w *worker) receive() {
	for {
		data, err := w.handler.Read()
		if err == nil {
			w.receiveChannel <- data
		}
	}
}

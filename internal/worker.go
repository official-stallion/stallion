package internal

import (
	"log"
	"net"
)

type worker struct {
	sendChannel    chan []byte
	receiveChannel chan []byte

	http http
	id   int
}

func NewWorker(id int, conn net.Conn) *worker {
	return &worker{
		http: http{
			conn: conn,
		},
		id: id,
	}
}

func (w *worker) Start() {
	defer close(w.sendChannel)

	go w.receive()

	select {
	case data := <-w.sendChannel:
		w.send(data)
	}
}

func (w *worker) send(data []byte) {
	if err := w.http.Write(data); err != nil {
		log.Fatalf("[%d] failed to send: %v\n", w.id, err)
	}
}

func (w *worker) receive() {
	for {
		data, err := w.http.Read()
		if err == nil {
			w.receiveChannel <- data
		}
	}
}

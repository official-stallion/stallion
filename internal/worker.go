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
	sendChannel chan []byte
	// receive channel is used for sending data to broker (its public)
	receiveChannel chan []byte
	// status channel
	statusChannel chan int
}

// newWorker generates a new worker.
// id: for worker id.
// conn: http connection over TCP.
// sen: sending channel.
// rec: receive channel.
func newWorker(id int, conn net.Conn, sen, rec chan []byte, sts chan int) *worker {
	return &worker{
		id: id,
		network: network{
			connection: conn,
		},

		sendChannel:    sen,
		receiveChannel: rec,
		statusChannel:  sts,
	}
}

// Start will start our worker.
func (w *worker) start() {
	// start for input data
	go w.arrival()

	// wait for any interrupt in send channel
	for {
		select {
		case data := <-w.sendChannel:
			w.transfer(data)
		}
	}
}

// transfer will send a data byte through handler.
func (w *worker) transfer(data []byte) {
	err := w.network.send(EncodeMessage(NewMessage(1, data)))
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

		m, _ := DecodeMessage(buffer)

		// passing data to broker channel
		w.receiveChannel <- m.Data
	}

	// announcing that the worker is done
	w.statusChannel <- w.id
}

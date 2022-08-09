package internal

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"
)

// worker handles a single client that wants
// to either subscribe or publish.
type worker struct {
	id         int
	connection net.Conn

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
		id:         id,
		connection: conn,

		sendChannel:    sen,
		receiveChannel: rec,
		statusChannel:  sts,
	}
}

// Start will start our worker.
func (w *worker) start() {
	// start for input data
	go w.receive()

	// wait for any interrupt in send channel
	for {
		select {
		case data := <-w.sendChannel:
			w.send(data)
		}
	}
}

// send will send a data byte through handler.
func (w *worker) send(data []byte) {
	writer := bufio.NewWriter(w.connection)
	if _, err := writer.Write(data); err != nil {
		log.Printf("failed to send: %v\n", err)
	}

	_ = writer.Flush()
	time.Sleep(10 * time.Millisecond)
}

// receive will check for input data from client.
func (w *worker) receive() {
	tmp := make([]byte, 1024)
	for {
		n, err := w.connection.Read(tmp)
		if err != nil {
			if err != io.EOF {
				log.Printf("read error: %s\n", err)
			}

			break
		}

		// passing data to broker channel
		w.receiveChannel <- tmp[:n]
	}

	// announcing that the worker is done
	w.statusChannel <- w.id
}

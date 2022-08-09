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
	statusChannel chan WorkChan
}

// newWorker generates a new worker.
// id: for worker id.
// conn: http connection over TCP.
// sen: sending channel.
// rec: receive channel.
func newWorker(id int, conn net.Conn, sen, rec chan []byte, sts chan WorkChan) *worker {
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
	err := w.network.send(encodeMessage(newMessage(Message, data)))
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
			log.Printf("json error: %v", er)

			continue
		}

		switch m.Type {
		case Message:
			// passing data to broker channel
			w.receiveChannel <- []byte(m.Data)
		case Subscribe:
			// passing subscribe message
			w.statusChannel <- WorkChan{
				id:      w.id,
				status:  true,
				channel: w.sendChannel,
			}
		case Unsubscribe:
			// passing unsubscribe message
			w.statusChannel <- WorkChan{
				id:     w.id,
				status: false,
			}
		}
	}

	// announcing that the worker is done
	w.statusChannel <- WorkChan{
		id:     w.id,
		status: false,
	}
}

package internal

import "net"

type worker struct {
	channel chan []byte
	http    http
}

func NewWorker(conn net.Conn) *worker {
	return &worker{
		http: http{
			conn: conn,
		},
	}
}

func (w *worker) start() {
	select {
	case data := <-w.channel:
		w.send(data)
	}
}

func (w *worker) send(data []byte) error {
	return w.http.Write(data)
}

func (w *worker) receive() []byte {
	data, err := w.http.Read()
	if err != nil {
		return nil
	}

	return data
}

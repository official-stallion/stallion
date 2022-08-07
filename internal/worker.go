package internal

type worker struct {
	channel chan []byte
}

func NewWorker() *worker {
	return &worker{}
}

func (w *worker) start() {
	select {
	case data := <-w.channel:
		w.send(data)
	}
}

func (w *worker) send(data []byte) {

}

func (w *worker) receive(data []byte) {

}

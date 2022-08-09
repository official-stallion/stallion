package internal

// workChan is worker channel with its id.
type workChan struct {
	id      int
	channel chan []byte
}

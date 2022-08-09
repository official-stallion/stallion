package internal

// workChan is worker channel with its id.
type workChan struct {
	id      int
	channel chan []byte
}

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

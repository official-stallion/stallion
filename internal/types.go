package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// WorkChan is worker channel with its id.
type WorkChan struct {
	id      int
	status  bool
	channel chan []byte
}

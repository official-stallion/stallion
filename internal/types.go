package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// WorkerChannel is worker channel with its id.
type WorkerChannel struct {
	id      int
	status  int
	topic   string
	channel chan Message
}

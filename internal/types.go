package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// WorkerChannel is worker channel with its id.
type workerChannel struct {
	id      int
	channel chan message
}

// SubscribeChannel is for subscribe data channel.
type subscribeChannel struct {
	id      int
	topic   string
	channel chan message
}

// UnsubscribeChannel is for unsubscribe data channel.
type unsubscribeChannel struct {
	id    int
	topic string
}

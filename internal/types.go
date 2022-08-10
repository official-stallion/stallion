package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// WorkerChannel is worker channel with its id.
type WorkerChannel struct {
	id      int
	channel chan Message
}

// SubscribeChannel is for subscribe data channel.
type SubscribeChannel struct {
	id      int
	topic   string
	channel chan Message
}

// UnsubscribeChannel is for unsubscribe data channel.
type UnsubscribeChannel struct {
	id    int
	topic string
}

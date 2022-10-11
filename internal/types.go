package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// workerChannel is worker channel with its id.
type workerChannel struct {
	id      int
	channel chan message
}

// subscribeChannel is for subscribe data channel.
type subscribeChannel struct {
	id      int
	topic   string
	channel chan message
}

// unsubscribeChannel is for unsubscribe data channel.
type unsubscribeChannel struct {
	id    int
	topic string
}

// pingMessage is the first message that is being sent
// to stallion server.
type pingMessage struct {
	username string
	password string
}

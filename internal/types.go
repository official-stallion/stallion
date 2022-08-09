package internal

// MessageHandler is a handler for messages that come from subscribing.
type MessageHandler func([]byte)

// Message is what we send between worker and clients.
type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// workChan is worker channel with its id.
type workChan struct {
	id      int
	channel chan []byte
}

package internal

// constant values for message types
const (
	Text        int = iota + 1 // normal message
	Subscribe                  // subscribe message
	Unsubscribe                // unsubscribe message
	PingMessage                // ping message
	PongMessage                // pong message
	Imposter                   // unauthorized user message
)

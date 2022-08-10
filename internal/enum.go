package internal

// constant values for message types
const (
	Text        int = iota + 1 // normal message
	Subscribe                  // subscribe message
	Unsubscribe                // unsubscribe message
)

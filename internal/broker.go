package internal

type arrChan []chan []byte

// broker handles the message sending and receiving.
type broker struct {
	channels       arrChan
	receiveChannel chan []byte
}

// NewBroker generates a broker.
func NewBroker() *broker {
	return &broker{}
}

// Start will start our broker logic.
func (b *broker) Start() {
	select {
	case data := <-b.receiveChannel:
		b.send(data)
	}
}

// Add will add subscribers to our broker.
func (b *broker) Add(channel chan []byte) {
	b.channels = append(b.channels, channel)
}

func (b *broker) send(data []byte) {
	for _, channel := range b.channels {
		channel <- data
	}
}

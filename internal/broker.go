package internal

type arrChan []chan []byte

type broker struct {
	channels map[string]arrChan
}

func NewBroker() *broker {
	return &broker{
		channels: make(map[string]arrChan),
	}
}

func (b *broker) send(data []byte, topic string) {
	for _, channel := range b.channels[topic] {
		channel <- data
	}
}

func (b *broker) add(channel chan []byte, topic string) {
	if _, ok := b.channels[topic]; ok {
		b.channels[topic] = append(b.channels[topic], channel)
	}
}

package main

import (
	"fmt"
	"sync"
)

// Message represents a message published by a publisher
type Message struct {
	Topic string
	Data  interface{}
}

// Subscriber represents a subscriber that receives messages
type Subscriber struct {
	ID      int
	Channel chan interface{}
}

// Update receives a message from the broker and handles it
func (s *Subscriber) Update() {
	for message := range s.Channel {
		fmt.Printf("Subscriber %d received message: %v\n", s.ID, message)
	}
}

// Broker represents the central message broker
type Broker struct {
	subscribers map[string][]*Subscriber
	mutex       sync.Mutex
}

// NewBroker creates a new instance of the Broker
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]*Subscriber),
	}
}

// Subscribe adds a subscriber to a specific topic
func (b *Broker) Subscribe(topic string, subscriber *Subscriber) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], subscriber)
}

// Unsubscribe removes a subscriber from a specific topic
func (b *Broker) Unsubscribe(topic string, subscriber *Subscriber) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	subscribers, ok := b.subscribers[topic]
	if !ok {
		return
	}
	for i, s := range subscribers {
		if s == subscriber {
			b.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			return
		}
	}
}

// Publish sends a message to all subscribers of a specific topic
func (b *Broker) Publish(msg Message) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	subscribers, ok := b.subscribers[msg.Topic]
	if !ok {
		return
	}
	for _, subscriber := range subscribers {
		subscriber.Channel <- msg.Data
	}
}

func main() {
	broker := NewBroker()
	defer closeBrokerChannels(broker)

	subscriber1 := &Subscriber{ID: 1, Channel: make(chan interface{})}
	subscriber2 := &Subscriber{ID: 2, Channel: make(chan interface{})}

	broker.Subscribe("topic1", subscriber1)
	broker.Subscribe("topic2", subscriber2)

	go subscriber1.Update()
	go subscriber2.Update()

	broker.Publish(Message{Topic: "topic1", Data: "Hello, World!"})
	broker.Publish(Message{Topic: "topic2", Data: "Greetings!"})

	// Output:
	// Subscriber 1 received message: Hello, World!
	// Subscriber 2 received message: Greetings!
}

// closeBrokerChannels closes the channels of all subscribers
func closeBrokerChannels(b *Broker) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for _, subscribers := range b.subscribers {
		for _, s := range subscribers {
			close(s.Channel)
		}
	}
}

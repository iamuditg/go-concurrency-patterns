package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Content string
}

type Subscriber struct {
	ID     int
	Events chan Message
	Done   chan struct{}
}

type Publisher struct {
	Subscribers map[int]Subscriber
	mu          sync.Mutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		Subscribers: make(map[int]Subscriber),
	}
}

func (p *Publisher) AddSubscriber(subscriber Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Subscribers[subscriber.ID] = subscriber
}

func (p *Publisher) RemoveSubscriber(subscriberID int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	subscriber, ok := p.Subscribers[subscriberID]
	if !ok {
		return
	}

	close(subscriber.Done)
	delete(p.Subscribers, subscriberID)
}

func (p *Publisher) PublishMessage(message Message) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, subscriber := range p.Subscribers {
		select {
		case subscriber.Events <- message:
		case <-subscriber.Done:
		}
	}
}

func main() {
	// Create a wait group to synchronize the completion of the goroutines
	var wg sync.WaitGroup

	// Create a publisher
	publisher := NewPublisher()

	// Create subscriber 1
	subscriber1 := Subscriber{
		ID:     1,
		Events: make(chan Message),
		Done:   make(chan struct{}),
	}
	publisher.AddSubscriber(subscriber1)

	// Create subscriber 2
	subscriber2 := Subscriber{
		ID:     2,
		Events: make(chan Message),
		Done:   make(chan struct{}),
	}
	publisher.AddSubscriber(subscriber2)

	// Launch goroutines for subscribers
	wg.Add(2)
	go subscriberRoutine(&subscriber1, &wg)
	go subscriberRoutine(&subscriber2, &wg)

	// Publish messages
	publisher.PublishMessage(Message{Content: "Hello, World!"})
	publisher.PublishMessage(Message{Content: "How are you?"})

	// Remove subscribers
	publisher.RemoveSubscriber(subscriber1.ID)
	publisher.RemoveSubscriber(subscriber2.ID)

	// Wait for all goroutines to complete
	wg.Wait()
}

func subscriberRoutine(subscriber *Subscriber, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case message, ok := <-subscriber.Events:
			if ok {
				fmt.Printf("Subscriber %d: Received message: %s\n", subscriber.ID, message.Content)
			} else {
				fmt.Printf("Subscriber %d: Shutting down\n", subscriber.ID)
				return
			}
		case <-subscriber.Done:
			fmt.Printf("Subscriber %d: Shutting down\n", subscriber.ID)
			return
		}
	}
}

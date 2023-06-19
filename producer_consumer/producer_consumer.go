package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Producer func(chan<- int)
type Consumer func(<-chan int)

type ProducerConsumer struct {
	buffer    chan int
	producer  Producer
	consumer  Consumer
	waitGroup *sync.WaitGroup
}

func NewProducerConsumer(bufferSize, numProducers, numConsumers int, producer Producer, consumer Consumer) *ProducerConsumer {
	return &ProducerConsumer{
		buffer:    make(chan int, bufferSize),
		producer:  producer,
		consumer:  consumer,
		waitGroup: &sync.WaitGroup{},
	}
}

func (pc *ProducerConsumer) Start() {
	for i := 0; i < cap(pc.buffer); i++ {
		pc.waitGroup.Add(1)
		go func() {
			defer pc.waitGroup.Done()
			pc.consumer(pc.buffer)
		}()
	}

	for i := 0; i < cap(pc.buffer); i++ {
		pc.waitGroup.Add(1)
		go func() {
			defer pc.waitGroup.Done()
			pc.producer(pc.buffer)
		}()
	}
}

func (pc *ProducerConsumer) Wait() {
	pc.waitGroup.Wait()
	close(pc.buffer)
}

// Example

func ExampleProducer(ch chan<- int) {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		data := rand.Intn(100)
		ch <- data
		fmt.Println("Produce:", data)
	}
}
func ExampleConsumer(ch <-chan int) {
	for data := range ch {
		fmt.Println("Consumed:", data)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func main() {
	producerConsumer := NewProducerConsumer(5, 2, 2, ExampleProducer, ExampleConsumer)
	producerConsumer.Start()
	producerConsumer.Wait()
}

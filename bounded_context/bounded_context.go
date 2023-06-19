package main

import (
	"fmt"
	"sync"
)

// ContextA represents the first bounded context
type ContextA struct {
	ID        int
	Input     chan string
	Output    chan string
	ShutDown  chan struct{}
	waitGroup *sync.WaitGroup
}

// ContextB represents the second bounded context
type ContextB struct {
	ID        int
	Input     chan string
	Output    chan string
	ShutDown  chan struct{}
	waitGroup *sync.WaitGroup
}

// NewContextA creates a new instance of ContextA
func NewContextA(id int, wg *sync.WaitGroup) *ContextA {
	return &ContextA{
		ID:        id,
		Input:     make(chan string),
		Output:    make(chan string),
		ShutDown:  make(chan struct{}),
		waitGroup: wg,
	}
}

// NewContextB creates a new instance of ContextB
func NewContextB(id int, wg *sync.WaitGroup) *ContextB {
	return &ContextB{
		ID:        id,
		Input:     make(chan string),
		Output:    make(chan string),
		ShutDown:  make(chan struct{}),
		waitGroup: wg,
	}
}

// Start starts the execution of ContextA
func (ctx *ContextA) Start() {
	defer ctx.waitGroup.Done()

	for {
		select {
		case message := <-ctx.Input:
			fmt.Printf("ContextA %d received: %s\n", ctx.ID, message)
			ctx.Output <- "Processed by ContextA"
		case <-ctx.ShutDown:
			fmt.Printf("ContextA %d shutting down\n", ctx.ID)
			return
		}
	}
}

// Start starts the execution of ContextB
func (ctx *ContextB) Start() {
	defer ctx.waitGroup.Done()

	for {
		select {
		case message := <-ctx.Input:
			fmt.Printf("ContextB %d received: %s\n", ctx.ID, message)
			ctx.Output <- "Processed by ContextB"
		case <-ctx.ShutDown:
			fmt.Printf("ContextB %d shutting down\n", ctx.ID)
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Create ContextA and ContextB
	contextA := NewContextA(1, &wg)
	contextB := NewContextB(2, &wg)

	// Start ContextA and ContextB
	wg.Add(2)
	go contextA.Start()
	go contextB.Start()

	// Send messages from ContextA to ContextB
	contextA.Input <- "Hello from ContextA"
	response := <-contextA.Output
	fmt.Println("Response from ContextA:", response)

	// Send messages from ContextB to ContextA
	contextB.Input <- "Hello from ContextB"
	response = <-contextB.Output
	fmt.Println("Response from ContextB:", response)

	// Shut down ContextA and ContextB
	close(contextA.ShutDown)
	close(contextB.ShutDown)

	// Wait for ContextA and ContextB to finish
	wg.Wait()

	fmt.Println("All contexts have shut down")
}

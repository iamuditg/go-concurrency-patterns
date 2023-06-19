package main

import (
	"fmt"
	"sync"
	"time"
)

// Barrier represents a synchronization point for a group of goroutines.
type Barrier struct {
	count  int
	wg     sync.WaitGroup
	mu     sync.Mutex
	notify chan struct{}
}

// NewBarrier creates a new Barrier with the specified count.
func NewBarrier(count int) *Barrier {
	return &Barrier{
		count:  count,
		notify: make(chan struct{}),
	}
}

// Wait waits for all goroutines to reach the barrier.
func (b *Barrier) Wait() {
	b.mu.Lock()
	b.count--
	if b.count == 0 {
		close(b.notify) // Notify all waiting goroutines
	}
	b.mu.Unlock()

	// Wait for the notification
	<-b.notify
}

func main() {
	const numWorkers = 5

	// Create a barrier with the number of expected workers
	barrier := NewBarrier(numWorkers)

	// Start multiple workers
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			fmt.Println("Worker", id, "started")
			// Simulate work
			time.Sleep(time.Second * time.Duration(id+1))
			fmt.Println("Worker", id, "finished work")

			// Wait for other workers to complete
			barrier.Wait()

			// Continue with the next phase of work
			fmt.Println("Worker", id, "continued with the next phase")
		}(i)
	}

	// Wait for all workers to complete
	barrier.Wait()

	fmt.Println("All workers completed")
}

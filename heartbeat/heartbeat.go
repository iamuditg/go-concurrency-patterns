package main

import (
	"fmt"
	"sync"
	"time"
)

// Heartbeat represents a heartbeat signal.
type Heartbeat struct {
	ID       int
	Interval time.Duration
}

// Worker represents a worker that sends heartbeat signals.
type Worker struct {
	ID        int
	Heartbeat chan Heartbeat
	Quit      chan bool
}

// NewWorker creates a new worker with the specified ID.
func NewWorker(id int) *Worker {
	return &Worker{
		ID:        id,
		Heartbeat: make(chan Heartbeat),
		Quit:      make(chan bool),
	}
}

// Start starts the worker to send heartbeat signals.
func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		for {
			select {
			case hb := <-w.Heartbeat:
				fmt.Printf("Worker %d received heartbeat %d\n", w.ID, hb.ID)
				time.Sleep(hb.Interval)
				fmt.Printf("Worker %d sent heartbeat %d\n", w.ID, hb.ID)
			case <-w.Quit:
				fmt.Printf("Worker %d quitting\n", w.ID)
				return
			}
		}
	}()
}

func main() {
	const numWorkers = 3

	var wg sync.WaitGroup

	// Create and start workers
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i)
		wg.Add(1)
		worker.Start(&wg)

		// Send heartbeat signals to workers
		go func(id int) {
			for j := 1; j <= 5; j++ {
				hb := Heartbeat{
					ID:       j,
					Interval: time.Second,
				}
				worker.Heartbeat <- hb
				time.Sleep(2 * time.Second)
			}

			worker.Quit <- true
		}(i)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed")
}

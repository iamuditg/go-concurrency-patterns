package main

import (
	"fmt"
	"time"
)

// Future represents a future result of a computation.
type Future struct {
	resultChan chan int
}

// Compute performs a computationally expensive operation asynchronously and return a Future.
func Compute() *Future {
	future := &Future{resultChan: make(chan int)}

	// Perform the computation asynchronously
	go func() {
		// Simulating a computationally expensive operation
		time.Sleep(time.Second * 3)

		// Set the result
		result := 42
		future.resultChan <- result
	}()

	return future
}

func (f *Future) GetResult() int {
	return <-f.resultChan
}

func main() {
	// Start the computation
	future := Compute()

	// Do other work concurrently
	//time.Sleep(40000)

	// Get the result when needed
	result := future.GetResult()
	fmt.Println("Result: ", result)
}

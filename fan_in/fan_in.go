package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels for input and output
	inputCh := make(chan int)
	outputCh := make(chan int)

	// Launch the producer goroutine
	go producer(input, inputCh)

	// Number of workers to process the input
	numWorkers := 3

	// Create a wait group to synchronize the completion of all workers
	var wg sync.WaitGroup

	// Launch the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(inputCh, outputCh, &wg)
	}

	// Close the output channel once all workers are done
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	// Process output values
	for val := range outputCh {
		fmt.Println(val)
	}
}

func producer(input []int, inputCh chan<- int) {
	defer close(inputCh)

	for _, val := range input {
		inputCh <- val
	}
}

func worker(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range inputCh {
		// Simulate some processing time
		time.Sleep(500 * time.Millisecond)

		// Process the value
		result := val * 2

		// Send the result to the output channel
		outputCh <- result
	}
}

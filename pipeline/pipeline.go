package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create channels for each stage of the pipeline
	inputCh := make(chan int)
	stage1Ch := make(chan int)
	stage2Ch := make(chan int)

	// Create a wait group to synchronize the completion of the stages
	var wg sync.WaitGroup

	// Launch the stages
	wg.Add(3)
	go stage1(inputCh, stage1Ch, &wg)
	go stage2(stage1Ch, stage2Ch, &wg)
	go stage3(stage2Ch, &wg)

	// Send input values to the input channel
	go func() {
		for _, val := range input {
			inputCh <- val
		}
		close(inputCh)
	}()

	// Wait for all stages to complete
	wg.Wait()
}

func stage1(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range inputCh {
		// Process stage 1 operation
		result := val * 2
		// Send the result to the next stage
		outputCh <- result
	}
	close(outputCh)
}

func stage2(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range inputCh {
		// Process stage 2 operation
		result := val + 1

		// Send the result to the next stage
		outputCh <- result
	}

	close(outputCh)
}

func stage3(inputCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range inputCh {
		// Process stage 3 operation
		result := val - 1

		// Output the final result
		fmt.Println(result)
	}
}

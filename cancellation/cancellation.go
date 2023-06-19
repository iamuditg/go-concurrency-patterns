package main

import (
	"context"
	"fmt"
	"time"
)

func doTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task canceled")
			return
		default:
			// Perform the task
			fmt.Println("Performing the task...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Start the task in a separate goroutine
	go doTask(ctx)

	// Wait for a few seconds
	time.Sleep(3 * time.Second)

	// Cancel the task
	cancel()

	// Wait for the task to finish
	time.Sleep(1 * time.Second)

	fmt.Println("Main goroutine completed")
}

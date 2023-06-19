package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a new context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Run a function concurrently with the context
	go func() {
		err := longRunningTask(ctx)
		if err != nil {
			fmt.Println("Task failed:", err)
		}
	}()

	// Wait for the context to be canceled or timeout
	select {
	case <-ctx.Done():
		fmt.Println("Context canceled")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout reached")
	}
}

func longRunningTask(ctx context.Context) error {
	// Simulate a long-running task
	time.Sleep(3 * time.Second)

	// Check if the context has been canceled
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Continue with the task
		fmt.Println("Task completed")
		return nil
	}
}

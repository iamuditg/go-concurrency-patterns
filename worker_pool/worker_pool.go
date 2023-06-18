package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Number of workers in the pool
	numWorkers := 3

	// Create a wait group to synchronize the completion of all workers
	var wg sync.WaitGroup

	// Create a channel for task distribution
	taskCh := make(chan Task)

	// Launch the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(taskCh, &wg)
	}

	// Send tasks to the worker pool
	tasks := generateTasks()

	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	for _, task := range tasks {
		taskCh <- task
	}

	close(taskCh)

	// Wait for all workers to complete
	wg.Wait()
}

// Task represents a unit of work to be performed
type Task struct {
	ID   int
	Data string
}

func worker(taskCh <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskCh {
		// Simulate task processing
		time.Sleep(500 * time.Millisecond)

		// Perform the task
		result := fmt.Sprintf("Task ID: %d, Data: %s", task.ID, task.Data)

		// Print the result
		fmt.Println(result)
	}
}

func generateTasks() []Task {
	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
		{ID: 4, Data: "Task 4"},
		{ID: 5, Data: "Task 5"},
	}

	return tasks
}

# Worker Pool Concurrency Pattern

## Description
The worker pool concurrency pattern is used to manage a fixed number of worker goroutines that concurrently execute tasks from a task queue or channel. The pool of workers allows for efficient handling of tasks in a concurrent and parallel manner. This pattern is especially useful when there is a large number of tasks that need to be processed concurrently.

## Implementation
In this example, we demonstrate the worker pool pattern by creating a fixed number of worker goroutines that process tasks from a task channel. The main steps involved in the implementation are as follows:

1. Create a task struct that represents a unit of work to be performed.
2. Create a channel for task distribution between the main goroutine and the worker goroutines.
3. Launch the desired number of worker goroutines, each waiting to receive tasks from the task channel.
4. Generate a set of tasks to be processed.
5. Send tasks to the task channel.
6. Each worker goroutine listens on the task channel for incoming tasks, processes them, and outputs the result.
7. Wait for all worker goroutines to complete processing.

## Usage
To run the example, execute the following command:
go run worker_pool.go
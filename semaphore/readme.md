# Semaphore in Go

## Description
The Semaphore is a synchronization primitive that allows controlling access to a shared resource. It maintains a count representing the number of available resources and provides two main operations: `Acquire` and `Release`. The `Acquire` operation requests to acquire a resource from the semaphore, blocking if none is currently available, and the `Release` operation signals that a resource is released and becomes available for other goroutines to acquire.

In this Go implementation of the Semaphore, a count-based semaphore is implemented using a mutex and a condition variable. The `NewSemaphore` function creates a new semaphore with the specified initial count, and the `Acquire` and `Release` methods are used to acquire and release resources, respectively.

## Usage
To use the Semaphore in your Go program, follow these steps:

1. Create a new semaphore using the `NewSemaphore` function, passing the initial count of available resources as an argument.
2. Call the `Acquire` method to acquire a resource from the semaphore. If no resources are available, the calling goroutine will be blocked until a resource becomes available.
3. Perform the necessary operations with the acquired resource.
4. Call the `Release` method to release the resource back to the semaphore, making it available for other goroutines to acquire.

## Example: Concurrent Task Execution with Semaphore
The following example demonstrates the usage of Semaphore by simulating concurrent task execution with limited resources.

### Task Function
The `ExampleTask` function represents a task that requires acquiring the semaphore to perform its work. In this example, the task simply prints a message indicating its progress.

### Main Function
In the `main` function, a semaphore instance is created using the `NewSemaphore` function with an initial count of available resources. Four goroutines are then started, representing concurrent tasks. Each task calls the `ExampleTask` function, acquires the semaphore, performs its work (represented by a sleep in this example), and releases the semaphore when finished. The program waits for all tasks to complete using a `sync.WaitGroup`.
go run semaphore.go

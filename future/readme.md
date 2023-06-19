# Future Concurrency Pattern in Go

The Future Concurrency pattern allows performing computationally expensive operations asynchronously and obtaining the result at a later point. It is useful when you want to initiate a long-running computation and continue executing other tasks concurrently without waiting for the result.

## Implementation
In this example, we have implemented the Future Concurrency pattern in Go. Here's how it works:

1. We define the Future struct, which represents a future result of a computation. It contains a resultChan channel to receive the result.
2. The Compute function performs a computationally expensive operation asynchronously and returns a Future. It creates a new goroutine that simulates the computation and sets the result on the resultChan channel.
3. The GetResult method of the Future waits for the computation to complete by receiving the result from the resultChan channel and returns the result.
4. In the main function, we start the computation by calling the Compute function, which returns a Future.
5. While the computation is running, we can perform other tasks concurrently.
6. When we need the result, we call the GetResult method on the Future to wait for the computation to complete and obtain the result.

## Usage
To use the Future Concurrency pattern in your Go programs:

1. Call the Compute function to initiate the computationally expensive operation asynchronously. It will return a Future representing the result.
2. Continue executing other tasks concurrently.
3. When you need the result, call the GetResult method on the Future to wait for the computation to complete and obtain the result.

The Future Concurrency pattern allows you to utilize the computational resources efficiently by performing expensive operations in the background while continuing with other tasks. It is especially useful when you have long-running computations that don't require immediate results.

Feel free to customize and adapt the code and readme according to your specific requirements and preferences for your Future Concurrency pattern implementation in Go.
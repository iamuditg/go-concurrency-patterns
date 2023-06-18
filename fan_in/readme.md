# Fan-In Concurrency Pattern

## Description
The fan-in concurrency pattern is used to combine the results of multiple concurrent workers into a single output. It involves multiple worker goroutines that independently process input values and produce their respective results. These results are then combined or merged into a single output channel. This pattern can help aggregate the results of parallel computations efficiently.

## Implementation
In this example, we demonstrate the fan-in pattern by processing a list of integers. We have a producer goroutine that sends the input values to a shared input channel. Multiple worker goroutines concurrently read from the input channel, perform some processing, and send their results to a shared output channel. The main goroutine reads from the output channel to retrieve the final combined results. The main steps involved in the implementation are as follows:

1. Create an input channel to receive the input values and an output channel to collect the processed results.
2. Launch a producer goroutine that sends the input values to the input channel.
3. Launch multiple worker goroutines. Each worker reads input values from the input channel, performs some processing, and sends the result to the output channel.
4. Use a wait group to synchronize the completion of all worker goroutines. Close the output channel once all workers are done processing.
5. Process the output values by reading from the output channel in the main goroutine.

## Usage
To run the example, execute the following command:
go run fan_in.go

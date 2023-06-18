# Pipeline Concurrency Pattern

## Description
The pipeline concurrency pattern is used to process a stream of data by breaking it down into multiple stages, each performing a specific operation on the data. Each stage consists of a goroutine that receives input from a channel, processes it, and sends the result to the next stage via another channel. This pattern enables parallel and efficient processing of data by dividing the workload into smaller, independent tasks.

## Implementation
In this example, we demonstrate the pipeline pattern by processing a stream of integers through multiple stages. Each stage performs a specific operation on the input and passes the transformed data to the next stage. The main steps involved in the implementation are as follows:

1. Create channels for each stage of the pipeline to facilitate the flow of data between stages.
2. Launch goroutines for each stage, passing the appropriate input and output channels.
3. Each stage goroutine performs its specific operation on the input data received from the input channel and sends the processed result to the output channel.
4. Use a wait group to synchronize the completion of all stages.
5. Send the initial input data to the first stage's input channel.
6. Wait for all stages to complete processing.

## Usage
To run the example, execute the following command:
go run pipeline.go
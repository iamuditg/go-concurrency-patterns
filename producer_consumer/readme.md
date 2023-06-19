# Producer-Consumer Pattern in Go

## Description
The Producer-Consumer pattern is a classic concurrency pattern where producers produce data and push it into a shared buffer, while consumers consume the data from the buffer. It decouples the producers and consumers, allowing them to work concurrently without direct coordination.

In this Go implementation of the Producer-Consumer pattern, a buffer (channel) is used as the shared resource between producers and consumers. Producers produce data and push it into the buffer, while consumers consume the data from the buffer. The number of producers and consumers can be specified, and they operate independently in separate goroutines.

## Usage
To use the Producer-Consumer pattern in your Go program, follow these steps:

1. Define a producer function that takes a send-only channel as an argument. This function will produce data and send it to the shared buffer.
2. Define a consumer function that takes a receive-only channel as an argument. This function will consume data from the shared buffer.
3. Create a ProducerConsumer instance using the `NewProducerConsumer` function, passing the buffer size, number of producers, number of consumers, the producer function, and the consumer function as arguments.
4. Start the ProducerConsumer instance by calling the `Start` method. This will start the producers and consumers in separate goroutines.
5. Optionally, you can wait for the producers and consumers to finish their work by calling the `Wait` method.

## Example: Producing and Consuming Data
The following example demonstrates the usage of the Producer-Consumer pattern by creating producers and consumers that produce and consume random integers.

### Producer
The `ExampleProducer` function is a producer that generates random integers and sends them to the shared buffer (channel). It simulates the production of data by sleeping for a random duration and sending the generated data to the buffer.

### Consumer
The `ExampleConsumer` function is a consumer that consumes data from the shared buffer (channel). It prints the consumed data and simulates the consumption process by sleeping for a random duration.

### Main Function
In the `main` function, a `ProducerConsumer` instance is created with a buffer size of 5, 2 producers, and 2 consumers. The `ExampleProducer` and `ExampleConsumer` functions are passed as arguments. The `Start` method is called to start the producers and consumers in separate goroutines, and the `Wait` method is called to wait for all producers and consumers to finish their work.

go run producer_consumer.go
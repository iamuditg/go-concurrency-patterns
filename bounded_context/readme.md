# Bounded-Context Pattern in Go
The Bounded-Context Pattern is used to isolate different parts of a system into bounded contexts, where each context is responsible for a specific functionality or domain. Communication between contexts is done through well-defined interfaces or channels. This pattern helps in managing complexity and ensuring clear separation of concerns in a concurrent system.

## Implementation
In this example, we have implemented the Bounded-Context Pattern in Go. We have simulated two bounded contexts, "ContextA" and "ContextB", communicating with each other through channels.

1. We define ContextA and ContextB structs to represent the two bounded contexts. Each context has an ID, input channel, output channel, shutdown channel, and a wait group to manage the goroutines.
2. The NewContextA and NewContextB functions are used to create instances of ContextA and ContextB respectively.
3. The Start methods of ContextA and ContextB are used to start the execution of the contexts. They listen for input on their input channels and process it accordingly. When a shutdown signal is received, the contexts gracefully shut down.
4. In the main function, we create instances of ContextA and ContextB, start their execution in separate goroutines, and demonstrate communication between them by sending messages and receiving responses.
5. Finally, we shut down the contexts by closing their shutdown channels and wait for them to finish using a wait group.

Feel free to customize and adapt the code and readme according to your specific requirements and preferences for your Bounded-Context Pattern implementation in Go.
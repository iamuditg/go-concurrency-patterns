# Cancellation Pattern in Go
The Cancellation Pattern is used to gracefully stop the execution of a concurrent operation. It allows a goroutine to signal its cancellation to other goroutines and terminate their execution.

## Implementation
In this example, we have implemented the Cancellation Pattern in Go using the context package.

1. We create a context with cancellation using context.WithCancel.
2. We define a function, doTask, that represents the concurrent task to be executed. It periodically checks for the cancellation signal using ctx.Done() and terminates its execution if the cancellation signal is received.
3. In the main function, we start the task in a separate goroutine and wait for a few seconds.
4. We cancel the task by calling the cancel function associated with the context.
5. Finally, we wait for the task to finish and print a completion message.

Feel free to customize and adapt the code and readme according to your specific requirements and preferences for your Cancellation Pattern implementation in Go.
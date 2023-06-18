# Rendezvous Concurrency Pattern

## Description
The Rendezvous concurrency pattern is used to synchronize two goroutines, ensuring that both goroutines reach a specific point before proceeding. It is commonly used when two goroutines need to exchange information or coordinate their actions.

## Implementation
In this example, we demonstrate the Rendezvous pattern by implementing two goroutines that need to synchronize their execution. The main steps involved in the implementation are as follows:

1. Create channels for the two goroutines to rendezvous.
2. Launch Goroutine 1 and Goroutine 2.
3. Implement the desired actions for each goroutine.
4. Signal the rendezvous point in Goroutine 1 and wait for Goroutine 2 to reach it.
5. Signal the rendezvous point in Goroutine 2 and wait for Goroutine 1 to reach it.

## Usage
To run the example, execute the following command:
go run rendezvous.go
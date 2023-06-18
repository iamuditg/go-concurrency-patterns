# Pub-Sub Concurrency Pattern

## Description
The Pub-Sub (Publisher-Subscriber) concurrency pattern is used for decoupling components in a system. Publishers publish events, and subscribers receive and process those events. It allows for a one-to-many relationship between publishers and subscribers, where multiple subscribers can receive events from one or more publishers.

## Implementation
In this example, we demonstrate the Pub-Sub pattern by implementing a simple publisher and two subscribers. The main steps involved in the implementation are as follows:

1. Create a `Publisher` struct that maintains a map of subscribers.
2. Create a `Subscriber` struct that contains an event channel for receiving events and a shutdown channel for graceful shutdown.
3. Implement methods for adding/removing subscribers and publishing events in the `Publisher` struct.
4. Launch goroutines for each subscriber to receive and process events.
5. Use channels for communication between the publisher and subscribers.
6. Publish events by sending them to the event channels of the subscribers.

## Usage
To run the example, execute the following command:
go run pubsub.go
# Pub-Sub Concurrency Pattern

## Description
This repository provides an implementation of the Publish-Subscribe (PubSub) concurrency pattern in Go, along with a Broker component. The PubSub pattern allows for loose coupling between publishers and subscribers, enabling scalable and flexible concurrent systems.

## Subscriber
The Subscriber struct represents a subscriber that receives messages from the broker. Each subscriber has a unique ID and a channel to receive messages.

## Broker
The Broker struct serves as the central message broker. It maintains a mapping of topics to subscribers. The broker allows subscribers to subscribe to specific topics and handles the distribution of messages to the appropriate subscribers. It provides methods to subscribe, unsubscribe, and publish messages.

## Message
The Message struct represents a message published by a publisher. It includes a topic and the actual data of the message.

## Usage
To use the PubSub concurrency pattern with the Broker component, follow these steps:

1. Create a new instance of the Broker using NewBroker().
2. Create one or more Subscriber instances and specify their unique IDs and channels for message reception.
3. Subscribe the subscribers to specific topics using the Subscribe() method of the broker.
4. Spawn goroutines for each subscriber and invoke their Update() methods to start receiving messages.
5. Publish messages to the topics of interest using the Publish() method of the broker.
6. Unsubscribe subscribers from topics when they are no longer interested in receiving messages using the Unsubscribe() method of the broker.
7. Ensure that all subscriber channels are properly closed when the program finishes executing.
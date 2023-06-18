package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a wait group to synchronize the completion of the goroutines
	var wg sync.WaitGroup

	// Create channels for the two goroutines to rendezvous and exchange data
	rendezvousCh1 := make(chan string)
	rendezvousCh2 := make(chan string)

	// Goroutine 1
	wg.Add(1)
	go goroutine1(rendezvousCh1, rendezvousCh2, &wg)

	// Goroutine 2
	wg.Add(1)
	go goroutine2(rendezvousCh1, rendezvousCh2, &wg)

	// Wait for both goroutines to complete
	wg.Wait()
}

// Goroutine 1
func goroutine1(rendezvousCh1 chan<- string, rendezvousCh2 <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	message := "Hello from Goroutine 1!"
	rendezvousCh1 <- message

	// Receive the response from Goroutine 2
	response := <-rendezvousCh2
	fmt.Println("Goroutine 1:", response)
}

// Goroutine 2
func goroutine2(rendezvousCh1 <-chan string, rendezvousCh2 chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Receive the message from Goroutine 1
	message := <-rendezvousCh1
	fmt.Println("Goroutine 2:", message)

	response := "Hello from Goroutine 2!"
	rendezvousCh2 <- response
}

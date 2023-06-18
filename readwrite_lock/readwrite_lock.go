package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Data struct {
	Value int
}

type Database struct {
	data   Data
	rwLock sync.RWMutex
}

func main() {
	// Create a wait group to synchronize the completion of the goroutines
	var wg sync.WaitGroup

	// Create a database instance
	db := Database{}

	// Launch multiple reader goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go reader(i+1, &db, &wg)
	}

	// Launch a writer goroutine
	wg.Add(1)
	go writer(&db, &wg)

	// Wait for all goroutines to complete
	wg.Wait()
}

// Reader goroutine
func reader(id int, db *Database, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire a read lock
	db.rwLock.RLock()

	// Simulate reading the data
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	// Read the data
	value := db.data.Value
	fmt.Printf("Reader %d: Read value %d \n\n", id, value)

	// Release the read lock
	db.rwLock.RUnlock()
}

// Writer goroutine
func writer(db *Database, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire a write lock
	db.rwLock.Lock()

	// Simulate updating the data
	time.Sleep(2 * time.Second)

	// Update the data
	db.data.Value = 42
	fmt.Println("Writer: Updated the data")

	// Release the write lock
	db.rwLock.Unlock()
}

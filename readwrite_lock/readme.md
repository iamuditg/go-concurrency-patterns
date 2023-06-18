# Read-Write Lock Concurrency Pattern

## Description
The Read-Write Lock concurrency pattern allows multiple readers to access a shared resource simultaneously, but only one writer can modify the resource at a time. This pattern is useful when the resource is predominantly read, and there is a need to ensure mutual exclusion when writing to the resource.

## Implementation
In this example, we demonstrate the Read-Write Lock pattern by implementing a simple database with a data structure that can be read by multiple reader goroutines and updated by a single writer goroutine. The main steps involved in the implementation are as follows:

1. Create a data structure (in this case, `Data`) that represents the shared resource.
2. Create a database (`Database`) that contains the shared data and a `sync.RWMutex` for synchronization.
3. Launch multiple reader goroutines that access and read the shared data.
4. Launch a writer goroutine that updates the shared data.
5. Acquire a read lock (`RLock()`) when reading the data and a write lock (`Lock()`) when updating the data.
6. Release the read lock (`RUnlock()`) and the write lock (`Unlock()`) when finished.

## Usage
To run the example, execute the following command:
go run readwrite_lock.go
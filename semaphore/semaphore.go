package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	count int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewSemaphore(initialCount int) *Semaphore {
	s := &Semaphore{count: initialCount}
	s.cond = sync.NewCond(&s.mutex)
	return s
}

func (s *Semaphore) Acquire() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for s.count <= 0 {
		s.cond.Wait()
	}
	s.count--
}

func (s *Semaphore) Release() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.count++
	s.cond.Signal()
}

func ExampleTask(id int, semaphore *Semaphore) {
	fmt.Printf("Task %d is waiting to acquire the semaphore\n", id)
	semaphore.Acquire()
	defer semaphore.Release()

	fmt.Printf("Task %d has acquired the semaphore\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Task %d has released the semaphore\n", id)
}

func main() {
	semaphore := NewSemaphore(2)

	var wg sync.WaitGroup
	wg.Add(4)

	for i := 1; i <= 4; i++ {
		go func(id int) {
			defer wg.Done()
			ExampleTask(id, semaphore)
		}(i)
	}

	wg.Wait()
}

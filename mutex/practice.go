package mutex

import (
	"fmt"
	"sync"
)

// Counter represents a shared counter with synchronized access.
type Counter struct {
	count int
	mu    sync.Mutex
}

// Increment increments the counter by 1, ensuring synchronized access.
func (c *Counter) Increment() {
	// Lock the mutex to ensure exclusive access to the counter
	c.mu.Lock()
	defer c.mu.Unlock()

	// Increment the counter by 1
	c.count++
}

func RunPractice() {
	// Initialize the Counter
	counter := Counter{}

	// Number of goroutines to create
	numGoroutines := 100

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Create goroutines to increment the counter concurrently
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Increment the counter
			counter.Increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final counter value:", counter.count)
}

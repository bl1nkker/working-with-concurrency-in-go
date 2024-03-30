// Assignment: Concurrent Summation

// Write a Go program that calculates the sum of numbers concurrently using multiple goroutines. Use a wait group to synchronize the completion of goroutines.

// Requirements:

// Define a function sum(numbers []int, result chan<- int, wg *sync.WaitGroup) that calculates the sum of numbers in a slice and sends the result to a channel. This function should take a slice of integers, a channel to send the result, and a pointer to a wait group.
// Inside the sum function, calculate the sum of numbers in the slice and send the result to the channel. Ensure that the wait group is notified upon completion of the function.
// In the main function, create a slice of integers containing some numbers.
// Initialize a channel to receive the result and a wait group.
// Launch multiple goroutines, each calling the sum function with a subset of numbers from the slice.
// Wait for all goroutines to finish using the wait group.
// Receive the results from the channel and calculate the total sum.
// Print the total sum.

package second

import (
	"fmt"
	"sync"
)

// sum calculates the sum of numbers in a slice and sends the result to a channel.
func sum(numbers []int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter when the function exits

	// Calculate the sum of numbers
	total := 0
	for _, num := range numbers {
		total += num
	}

	// Send the result to the channel
	result <- total
}

func RunAssignment() {
	// Slice of numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Channel to receive the result
	result := make(chan int)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Number of goroutines to create
	numGoroutines := 5
	wg.Add(numGoroutines) // Increment the WaitGroup counter

	// Divide the slice into subsets for each goroutine
	subsetSize := len(numbers) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * subsetSize
		end := start + subsetSize
		if i == numGoroutines-1 {
			end = len(numbers)
		}
		go sum(numbers[start:end], result, &wg) // Start goroutine for each subset
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the result channel after all calculations are done
	close(result)

	// Receive and accumulate the results
	totalSum := 0
	for res := range result {
		totalSum += res
	}

	// Print the total sum
	fmt.Println("Total sum:", totalSum)
}

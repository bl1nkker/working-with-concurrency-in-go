package producerconsumer

import (
	"fmt"
	"sync"
)

const NumWorkers = 3
const NumJobs = 10

func producer(jobs chan<- int) {
	for i := 0; i < NumJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func RunExample() {
	fmt.Println("Starting producer-consumer...")

	// Create a channel to communicate between producer and consumers
	jobs := make(chan int)

	// Create a wait group to wait for all consumers to finish
	var wg sync.WaitGroup

	// Start consumers
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}

	// Start producer
	go producer(jobs)

	// Wait for all consumers to finish
	wg.Wait()

	fmt.Println("All jobs processed.")
}

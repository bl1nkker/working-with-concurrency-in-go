package producerconsumer

import (
	"fmt"
	"time"
)

func RunSelect() {
	// Create two channels
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Start a goroutine to send data on ch1 after 1 second
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	// Start another goroutine to send data on ch2 after 2 seconds
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "hello"
	}()

	// Use select to wait for data on either channel
	select {
	case num := <-ch1:
		fmt.Println("Received from ch1:", num)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		fmt.Println("Timeout: No data received")
	}
}

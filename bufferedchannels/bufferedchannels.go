package bufferedchannels

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got i from channel", i)
		// simulate doing a lot of work
		time.Sleep(5 * time.Second)
	}
}

func Run() {
    // Buffered channel. By default it is set to 1
	ch := make(chan int, 10)
	go listenToChan(ch)
	for i := 0; i <= 100; i++ {
		fmt.Println("Sending i to chan:", i)
		ch <- i
		fmt.Println("Sent i to chan:", i)
	}
	fmt.Println("Done")
	close(ch)
}

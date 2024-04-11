package simplechannels

import (
	"fmt"
	"strings"
)

// "<-" - This is a receive-only channel of type string. It means that the function can only receive values from this channel.
// "->" - This is a send-only channel of type string. It means that the function can only send values to this channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		fmt.Println("Shout started...")
		// When you get something from the channel ping, put it inside this variable
		sPing, ok := <-ping

		if !ok{
			// do something here
		}
		fmt.Println("Received value from ping...")
		// Sending back to the channel pong
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(sPing))
		fmt.Println("Sent value to pong...")
	}
}

func Run() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER")
	for {
		fmt.Print("-> ")
		// get user input
		var userInput string
		_, err := fmt.Scanln(&userInput)

		if err != nil {
			// Error: type not a pointer: string
			fmt.Println("Error:", err)
		}

		if userInput == "Q" {
			break
		}
		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All Done. Closing channels!")
	close(ping)
	close(pong)
}

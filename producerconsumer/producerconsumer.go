package producerconsumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const PizzasAmount = 10

var pizzasMade, pizzasFailed, total int

type Producer struct{
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct{
	pizzaNumber int
	message string
	success bool
}

// Close is a method for gracefully shutting down the Producer.
// It sends a signal to stop processing and waits for confirmation.
func (p *Producer) Close() error {
    // Create a new channel to receive an error signal
    ch := make(chan error)
    
    // Send the channel into the quit channel of the Producer
    p.quit <- ch
    
    // Wait to receive an error signal from the channel
    // indicating whether the shutdown was successful
    return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder{
	pizzaNumber++
	if pizzaNumber <= PizzasAmount{
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		message := ""
		success := false
		
		fmt.Printf("Making pizza #%d. It will take %d seconds\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2{
			pizzasFailed ++
			message = fmt.Sprintf("*** We ran out of ingrеdients for pizza #%d!\n", pizzaNumber)
		} else if rnd <= 4{
			pizzasFailed ++
			message = fmt.Sprintf("*** The cook quit while making the pizza #%d!\n", pizzaNumber)
		} else{
			pizzasMade ++
			success = true
			message = fmt.Sprintf("Pizza order #%d is ready!\n", pizzaNumber)
		}
		total ++
		return &PizzaOrder{pizzaNumber: pizzaNumber, message: message, success: success}
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer){
	// keep track of which pizza we are making
	i := 0
	// run forever, or until we receive a quit notification
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil{
			i = currentPizza.pizzaNumber
			// The select statement in Go selects cases based on which channel operation becomes ready first. If multiple cases are ready at the same time, one will be chosen randomly.
			select{
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func Run(){
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("------------------------------")

	// Create Producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		// Wtf?
		quit: make(chan chan error),
	}

	// Run producer in background
	go pizzeria(pizzaJob)

	// create and run the consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= PizzasAmount{
			if i.success{
				color.Green(i.message)
				color.Green("Order number %d is out of delivery!", i.pizzaNumber)
			} else{
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else{
			color.Cyan("Done making pizzas.")
			err := pizzaJob.Close()
			if err != nil{
				color.Red("*** Error closing channel...", err)
			}
		}
	}
	color.Cyan("------------------------------")
	color.Cyan("Done for the day!")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)
	switch{
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an OK day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day!")
	}
	
}
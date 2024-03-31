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

		if rnd < 5{
			pizzasFailed ++
		} else{
			pizzasMade ++
		}
		total ++
		fmt.Printf("Making pizza #%d. It will take %d seconds", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2{
			message = fmt.Sprintf("*** We ran out of ingridients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4{
			message = fmt.Sprintf("*** The cook quit while making the pizza #%d!", pizzaNumber)
		} else{
			success = true
			message = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
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
	}
	// try to make pizzas
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
}
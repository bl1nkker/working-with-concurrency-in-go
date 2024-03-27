package producerconsumer

import "github.com/fatih/color"

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

func pizzeria(pizzaMaker *Producer){
	// keep track of which pizza we are making

	// run forever, or until we receive a quit notification

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
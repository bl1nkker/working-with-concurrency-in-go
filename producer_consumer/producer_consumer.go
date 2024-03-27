package producerconsumer

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

func Run(){

}
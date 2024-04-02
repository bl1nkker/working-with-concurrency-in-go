package dining

import (
	"fmt"
	"time"
)

// The Dining Philosophers Problem is a classic synchronization problem used to illustrate the challenges of ensuring the safe sharing
// of finite resources among multiple processes, especially in concurrent programming. The problem was formulated by Edsger Dijkstra in 1965.
// In the problem, there are five philosophers sitting around a circular table. Each philosopher alternates between thinking and eating.
// To eat, a philosopher needs two forks, one on the left and one on the right. However, there are only five forks available, one between each pair of adjacent philosophers.

// The challenge is to design a protocol (algorithm) that allows the philosophers to dine without experiencing deadlock
// (where each philosopher holds one fork and is waiting for the other) or starvation (where a philosopher is unable to acquire the necessary forks to eat).

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socratus", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 // How many times does the person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func Run() {
	// Print welcome message
	fmt.Println("Welcome, fuckers!")
	// Start the meal
	dine()
	// Print finished message
	fmt.Println("Bye, fuckers!")
}

func dine() {

}

package dining

import (
	"fmt"
	"sync"
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
	fmt.Println("Dining Philosopher Problem")
	fmt.Println("The table is empty")
	fmt.Println("--------------------------")
	// Start the meal
	dine()
	// Print finished message
	fmt.Println("--------------------------")
	fmt.Println("The table is empty")
}

func dine() {
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers)) // When this reaches 0, everyone is ready to begin eating
	fmt.Println("Initialized seated wait groups!")

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers)) // When this reaches 0, everyone is done eating
	fmt.Println("Initialized done wait groups!")

	// Mutexes here need to lock the "forks"
	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	fmt.Println("Initialized forks with mutexes!")

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off the goroutine for the current philosopher
		diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait() // pause program execution until all five goroutines done (until all five philosophers done eating)
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("--- Dining Problem started for: %v\n", philosopher)
}

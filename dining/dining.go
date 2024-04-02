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
var eatTime = 0 * time.Second
var thinkTime = 0 * time.Second
var sleepTime = 0 * time.Second

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
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait() // pause program execution until all goroutines done (until all philosophers done eating)
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	// Wait until all philosophers will seat
	seated.Wait()

	// eat {hunger} times
	for i := hunger; i > 0; i-- {
		// We need to use this, in case of logical race condition
		if philosopher.leftFork > philosopher.rightFork {
			// get a lock on both philosopher's forks (lock a mutex). It will not be blocked, if it is already blocked by another philosopher
			fmt.Printf("\t%s is trying to lock his left fork: %d...\n", philosopher.name, philosopher.leftFork)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork: %d.\n", philosopher.name, philosopher.leftFork)

			fmt.Printf("\t%s is trying to lock his right fork: %d...\n", philosopher.name, philosopher.rightFork)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork: %d.\n", philosopher.name, philosopher.rightFork)
		} else {
			// get a lock on both philosopher's forks (lock a mutex). It will not be blocked, if it is already blocked by another philosopher
			fmt.Printf("\t%s is trying to lock his right fork: %d...\n", philosopher.name, philosopher.rightFork)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork: %d.\n", philosopher.name, philosopher.rightFork)

			fmt.Printf("\t%s is trying to lock his left fork: %d...\n", philosopher.name, philosopher.leftFork)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork: %d.\n", philosopher.name, philosopher.leftFork)
		}

		fmt.Printf("\t%s has both forks and is eating...\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking...\n", philosopher.name)
		time.Sleep(thinkTime)

		fmt.Printf("\t+ %s is done eating and thinking. Freeing the forks %d and %d...\n", philosopher.name, philosopher.leftFork, philosopher.rightFork)
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t++ %s put down the forks. The forks %d and %d is Free!\n", philosopher.name, philosopher.leftFork, philosopher.rightFork)
	}

	fmt.Printf("%s is satisfied.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)
}

package second

import (
	"fmt"
	"sync"
	"time"
)

func Run(){
	var wg sync.WaitGroup
	words := []string{"Alpha", "Beta", "Delta", "Gamma", "Zeta", "Epsilon", "Theta", "Eta"}
	wg.Add(len(words))
	for i, val := range words{
		go PrintSomething(fmt.Sprintf("%d: %s", i, val), &wg)
	}
	// Without sleep goroutine above will not display anything
	wg.Wait()
	wg.Add(1)
	PrintSomething("This is the second thing to be printed", &wg)
}

func PrintSomething(s string, wg *sync.WaitGroup){
	// defer is used to schedule a function call to be executed when the surrounding function (the one containing the defer statement) returns, either normally or via a panic.
	defer wg.Done()
	fmt.Println(s)
}

func RunRestaurantSimulation() {
	var wg sync.WaitGroup

	// Number of tables in the restaurant
	numTables := 5

	// Number of customers per table
	numCustomers := 3

	// Add the number of tasks (tables) to the wait group
	wg.Add(numTables)

	// Serve each table concurrently
	for table := 1; table <= numTables; table++ {
		go func(table int) {
			defer wg.Done() // Decrease the counter when the table is served
			ServeTable(table, numCustomers)
		}(table)
	}

	// Wait for all tables to be served
	wg.Wait()

	// Additional cleanup task
	fmt.Println("Restaurant closed. Goodbye!")
}

func ServeTable(table int, numCustomers int) {
	fmt.Printf("Table %d seated and ready to order\n", table)

	// Simulate taking orders
	time.Sleep(1 * time.Second)

	fmt.Printf("Table %d orders taken\n", table)

	// Simulate cooking and serving food
	time.Sleep(2 * time.Second)

	fmt.Printf("Table %d food served\n", table)

	// Simulate customers eating
	time.Sleep(3 * time.Second)

	fmt.Printf("Table %d customers finished eating\n", table)

	// Simulate clearing the table
	time.Sleep(1 * time.Second)

	fmt.Printf("Table %d cleared\n", table)
}
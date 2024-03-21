package second

import (
	"fmt"
	"sync"
)

func RunSecond(){
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
	defer wg.Done()
	fmt.Println(s)
}
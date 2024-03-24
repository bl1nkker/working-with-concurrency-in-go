package mutex

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, wg *sync.WaitGroup){
	wg.Done()
	msg = s
}

func raceConditionFunc(){
	msg = "Hello World"
	wg.Add(2)
	go updateMessage("This should be First", &wg)
	go updateMessage("This should be Second", &wg)
	wg.Wait()
	fmt.Println(msg)
}

func Run(){
	raceConditionFunc()
}
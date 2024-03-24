package mutex

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex){
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func updateMessageRC(s string){
	defer wg.Done()
	msg = s
}

func RunRC(){
	msg = "Hello World"
	wg.Add(2)
	go updateMessageRC("This should be First")
	go updateMessageRC("This should be Second")
	wg.Wait()
	fmt.Println(msg)
}

func Run(){
	var mutex sync.Mutex
	msg = "Hello World"
	wg.Add(2)
	go updateMessage("This should be First", &mutex)
	go updateMessage("This should be Second", &mutex)
	wg.Wait()
	fmt.Println(msg)
}

package second

import (
	"fmt"
	"time"
)

func RunSecond(){
	go PrintSomething("This is the first thing to be printed")
	// Without sleep goroutine above will not display anything
	time.Sleep(1 * time.Second)
	PrintSomething("This is the second thing to be printed")
}

func PrintSomething(s string){
	fmt.Println(s)
}
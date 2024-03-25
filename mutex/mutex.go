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
	wg.Add(10)
	go updateMessageRC("This should be 1")
	go updateMessageRC("This should be 2")
	go updateMessageRC("This should be 3")
	go updateMessageRC("This should be 4")
	go updateMessageRC("This should be 5")
	go updateMessageRC("This should be 6")
	go updateMessageRC("This should be 7")
	go updateMessageRC("This should be 8")
	go updateMessageRC("This should be 9")
	go updateMessageRC("This should be 10")
	fmt.Println(msg)
}

func Run(){
	var mutex sync.Mutex
	msg = "Hello World"
	wg.Add(10)
	go updateMessage("This should be 1", &mutex)
	go updateMessage("This should be 2", &mutex)
	go updateMessage("This should be 3", &mutex)
	go updateMessage("This should be 4", &mutex)
	go updateMessage("This should be 5", &mutex)
	go updateMessage("This should be 6", &mutex)
	go updateMessage("This should be 7", &mutex)
	go updateMessage("This should be 8", &mutex)
	go updateMessage("This should be 9", &mutex)
	go updateMessage("This should be 10", &mutex)
	wg.Wait()
	fmt.Println(msg)
}

func RunClearExample(){
	var wg sync.WaitGroup
    var sharedNumber int
	var mutex sync.Mutex
	amount := 1000

    // Добавляем две горутины в WaitGroup
    wg.Add(amount)

    for i := 0; i < amount; i ++{
		go func() {
			defer wg.Done()
			sharedNumber++
		}()
	}

    // Ждем, пока обе горутины завершат выполнение
    wg.Wait()

    // Выводим результат
    fmt.Println("Shared number without mutex:", sharedNumber)
	sharedNumber = 0
	// Добавляем две горутины в WaitGroup
    wg.Add(amount)

	for i := 0; i < amount; i ++{
		go func() {
			defer wg.Done()
			mutex.Lock()
			sharedNumber++
			mutex.Unlock()
		}()
	}

    // Ждем, пока обе горутины завершат выполнение
    wg.Wait()

    // Выводим результат
    fmt.Println("Shared number with mutex:", sharedNumber)
}

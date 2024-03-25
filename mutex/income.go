package mutex

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

type Income struct{
	Source string
	Amount int
}

func RunIncome(){
	var balance int
	var mutex sync.Mutex

	fmt.Printf("Initial account balance is %d.00\n", balance)
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 100},
		{Source: "Freelance", Amount: 250},
		{Source: "Side Projects", Amount: 730},
	}

	wg2.Add(len(incomes))
	for i, income := range incomes{
		go func(i int, income Income){
			defer wg2.Done()
			for week := 1; week <= 52; week ++{
				mutex.Lock()
				balance += income.Amount
				mutex.Unlock()
				fmt.Printf("On week %d you earned %d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg2.Wait()
	fmt.Printf("Final bank balance: %d\n", balance)
}
package main

import (
	"working-with-concurrency-in-go/mutex"
	"working-with-concurrency-in-go/second"
)

func main(){
	// second.Run()
	// second.RunRestaurantSimulation()
	second.RunChallenge()
	// mutex.Run()
	// mutex.RunRC()
	// mutex.RunClearExample()
	mutex.RunIncome()
}
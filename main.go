package main

import (
	"working-with-concurrency-in-go/second"
	"working-with-concurrency-in-go/third"
)

func main(){
	second.Run()
	third.Run()
}
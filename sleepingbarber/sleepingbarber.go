// The Sleeping Barber Problem is a classic synchronization problem in computer science and concurrent programming.
// It illustrates challenges in managing resources among multiple processes or threads. The scenario involves a barber shop with a
// barber who sleeps when there are no customers and a waiting room with limited seating capacity. The problem is to design a system
// where customers arrive, wait in the waiting room if it's full, and are served by the barber one at a time. The challenge lies in
// coordinating the access to the barber and the waiting room to avoid race conditions and ensure that customers are served in a fair
// and orderly manner without deadlocks or starvation. It's often used to demonstrate the importance of synchronization mechanisms like
// semaphores, mutexes, or condition variables in concurrent programming.
// 1. The Barber's Behavior:
// - The barber falls asleep if there are no customers to serve.
// - When a customer arrives, the barber wakes up and serves the customer.
// 2. The Waiting Room:
// - The waiting room has a limited number of chairs, say, N.
// - If the waiting room is full when a customer arrives, the customer leaves.
// 3. Customer Behavior:
// - Customers arrive at random intervals.
// - If the barber is busy (serving another customer), a customer waits in the waiting room if there's space.
// - If the waiting room is full, the customer leaves.
// 4. Service Completion:
// - After a customer is served, they leave the barbershop.
// These rules highlight the coordination and synchronization required to manage the access to shared resources
// (the barber and the waiting room) among multiple customers and the barber. Implementing a solution to the Sleeping Barber Problem
// involves ensuring that these rules are followed to prevent issues like deadlocks, race conditions, or starvation.

package sleepingbarber

import (
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func Run() {
	// print welcome message
	color.Green("The sleeping barber problem")
	color.Green("--------------------------")

	// create channels if we need any
	// Maximum 10 people in the channel
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the Barbershop struct
	shop := Barbershop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            false,
	}

	color.Cyan("The shop is open for the day!")

	// add barbers
	shop.AddBarber("Frank")

	// start the barbershop as a go routiner
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		// this time.After(timeOpen) is blocking the closing code for "timeOpen" seconds.
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()
	// add clients

	// block until the barbershop is closed (forloop)
	time.Sleep(15 * time.Second)
}

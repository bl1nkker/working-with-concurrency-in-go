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

func Run() {
	// print welcome message

	// create channels if we need any

	// create the Barbershop struct

	// add barbers

	// start the barbershop as a go routiner

	// add clients

	// block until the barbershop is closed (forloop)
}

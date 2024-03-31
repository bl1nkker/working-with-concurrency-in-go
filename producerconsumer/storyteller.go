package producerconsumer

import (
	"fmt"
	"time"
)

type Parent struct {
	stories chan string
	// In Go, the chan struct{} type is commonly used as a signaling mechanism. Here's why:
	// Minimal Memory Overhead: Using an empty struct (struct{}) as the type for a channel indicates that we're only interested in the signaling aspect of the channel, not in passing any actual data. This results in minimal memory overhead because an empty struct consumes zero bytes of memory.
	// Clear Intent: When other developers see a channel with type chan struct{}, it clearly communicates that the channel is being used for signaling purposes only. There is no ambiguity about whether data is being transmitted through the channel.
	// Efficient Signaling: Since an empty struct consumes no memory, sending and receiving signals through this channel is very efficient. There's no data to copy or allocate.
	// Blocking and Synchronization: Sending a value through a channel of type chan struct{} blocks until another goroutine receives the value, making it a simple and effective way to synchronize between goroutines.
	stop    chan struct{}
}

func NewParent() *Parent {
	return &Parent{
		stories: make(chan string),
		stop:    make(chan struct{}),
	}
}

func (p *Parent) ReadStory(story string) {
	p.stories <- story
}

func (p *Parent) Bedtime() {
	// If you remove the line close(p.stories), it can potentially lead to a deadlock between your main goroutine and child goroutines.
	close(p.stories) // Close the stories channel to indicate bedtime
	conf := <-p.stop         // Wait for confirmation from the child
	fmt.Printf("Confirmation %s", conf)
	fmt.Println("Turning off the lights. Goodnight!")
}

type Child struct {
	stories <-chan string
}

func NewChild(parent *Parent) *Child {
	return &Child{
		stories: parent.stories,
	}
}

func (c *Child) Listen(parent *Parent) {
	for story := range c.stories {
		fmt.Println("Parent reads:", story)
	}
	fmt.Println("Child falls asleep.")
	time.Sleep(2 * time.Second) // Simulate child falling asleep
	parent.stop <- struct{}{}    // Signal parent that child is asleep
}

func RunStoryteller() {
	parent := NewParent()
	child := NewChild(parent)

	go child.Listen(parent) // Child listens to bedtime stories
	parent.ReadStory("Once upon a time...")
	parent.ReadStory("There was a little bunny...")
	parent.Bedtime() // Parent stops reading and waits for child to sleep

	// Wait for a few seconds to ensure the program doesn't terminate immediately
	time.Sleep(3 * time.Second)
}

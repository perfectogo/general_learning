package main

import (
	"fmt"
)

func sender(ch chan int) {
	fmt.Println("Sending data...")
	ch <- 42 // Send data to the channel
	fmt.Println("Data sent!")
}

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Start a goroutine to read data from the channel
	go func() {
		fmt.Println("Reading data...")
		data := <-ch // Read data from the channel
		fmt.Println("Received data:", data)
	}()

	// Pause for a moment to ensure the reader has started
	fmt.Println("Waiting for the reader to start...")
	for i := 0; i < 5; i++ {
		fmt.Println("Tick...")
	}

	// Send data to the channel
	sender(ch)

	fmt.Println("Program completed!")
}

package main

import (
	"context"
	"fmt"
	"time"
)

//select
// select allows a goroutine to wait for multiple channel operations and execute whichever operation is ready.

// Think of
// it like waiting for multiple phone calls. Whichever call arrives first, you respond to it.

// without

// Problem: If ch1 doesn't send anything, your program gets stuck waiting there, even if ch2 has data.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello from channel 1"
	}()

	go func() {
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg2 := <-ch2:
		fmt.Println(msg2)

	case msg1 := <-ch1:
		fmt.Println(msg1)
	}

	// 	Imagine:
	// - Goroutine A is waiting for data from a channel.
	// - The data might take a long time to arrive.
	// - Instead of waiting forever, we give it a timeout.

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "API response received"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)

		// This is particularly useful in backend development.
		// Imagine you call an external API. You don't want your application to wait forever.
		// Go provides time.After().
	case <-time.After(2 * time.Second):
		fmt.Println("Request timeout")
	}

	// there is a goroutine leak problem
	// worker sends its result after 5s
	// But main finishes : The program exits, so the worker is terminated along with the process.
	// If the application were still running: The worker would wake up after 5 seconds and get stuck trying to send data to ch, because nobody is receiving.

	// to solve goroutine problem use
	// 1. buffered channel even if nobody receives it, the channel can store that value.

	// better way to solve this problem is use context.withTimeout()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	select {
	case data := <-ch:
		fmt.Println(data)

	case <-ctx.Done():
		fmt.Println("Request cancelled:", ctx.Err())
	}
}

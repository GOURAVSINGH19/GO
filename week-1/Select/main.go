package main

import (
	"fmt"
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
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

package main

import (
	"fmt"
	"sync"
)

// go routines -> lightwait thread of execution managed by the GO runtime.

// It allow multiple function execute concurrently

// what is concurrently  means ?
// Concurrently means executing multiple tasks during overlapping periods of time, instead of completing one task entirely before starting another.

// example
// Imagine you are a chef:
// - Task 1: Boil water.
// - Task 2: Chop vegetables.
// - Task 3: Prepare sauce.
// You start boiling water, then chop vegetables while the water is boiling.
// You are making progress on multiple tasks concurrently.

// example :-
// For example, imagine a backend server handling 100 user requests.
// Instead of waiting for each request to finish before starting the next,
// Go can use goroutines to handle work concurrently.

// diff btw concurrently / Parallellism
// Concurrency is about dealing with multiple tasks at once.
// Parallelism is about doing multiple tasks at the same time.

func task(wg *sync.WaitGroup) {
	// Marks one task as complete
	// defer user for :- We use defer to postpone function execution until the surrounding function returns.
	// It is commonly used for cleanup operations such as closing files, releasing mutexes, closing database resources, and calling WaitGroup.Done().

	// Important
	// defer doesn't execute immediately. It schedules the function to execute when the surrounding function returns.
	defer wg.Done()
	fmt.Println("Task is running....")

	// this cause problem Suppose If func has multiple return points. then it returns without calling wg.Done(). So wg.Wait() might wait forever.
	// wg.Done()
}

// without go keyword
// The function executes synchronously: main() waits for task() to finish.
// with go keyword ( It create a new go thread)

// >>>>>>>>>>>>> channels
// A channel is a communication mechanism that allows goroutines to send and receive values.

// An unbuffered channel blocks the sender until a receiver is ready, and blocks the receiver until a sender is ready.

// what is unbuffered
// means when channel has 0 capacity
// Sender and receiver must synchronize.

// what is buffered
// means when channel has capacity
// Sender can proceed while buffer has space

func sendData(ch chan int) {
	// sending data
	ch <- 100
	ch <- 20

	fmt.Println("Data sent")

}

// This func only send int chan<- int
func send1(ch chan<- int) {
	// sending data
	ch <- 100
	ch <- 20

	fmt.Println("Data sent")

}

// This func only receive int <-chan int
func send2(ch <-chan int) {
	// sending data
	value := <-ch

	fmt.Println("Data Received", value)

}

// Only the sender should normally close a channel.
// Receiving from a closed, drained channel returns the zero value immediately.
// Sending to a closed channel causes a panic.
// Closing an already closed channel causes a panic.
// You don't have to close every channel.

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
	}
}
func main() {

	// var wg sync.WaitGroup

	// Increases the number of tasks to wait for
	// wg.Add(1)
	// go task(&wg)
	// // Blocks until the counter becomes zero
	// wg.Wait()
	// fmt.Println("Main function")

	// >>>>>>>>>>>>>>> Channels

	ch := make(chan int, 3)

	go sendData(ch)

	// receiving data
	// value := <-ch
	// value2 := <-ch

	// fmt.Println(value)
	// fmt.Println(value2)

	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 0; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

}

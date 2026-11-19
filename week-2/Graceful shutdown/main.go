package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	numWorkers      = 3
	numJobs         = 20
	queueSize       = 5
	jobDuration     = 1 * time.Second
	shutdownTimeout = 5 * time.Second // try 1 * time.Second to see in-flight jobs get cancelled
)

// worker processes jobs until the jobs channel is closed and drained,
// or until workCtx is cancelled (shutdown deadline exceeded).
func worker(workCtx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Deadline passed: don't start any more queued jobs.
		if workCtx.Err() != nil {
			fmt.Printf("Worker %d skipping Job %d: %v\n", id, job, workCtx.Err())
			continue
		}

		fmt.Printf("Worker %d processing Job %d\n", id, job)

		// Simulate job processing that can be interrupted.
		select {
		case <-time.After(jobDuration):
			fmt.Printf("Worker %d completed Job %d\n", id, job)
		case <-workCtx.Done():
			fmt.Printf("Worker %d cancelled Job %d: %v\n", id, job, workCtx.Err())
		}
	}

	fmt.Printf("Worker %d stopped\n", id)
}

// producer sends jobs until all are submitted or ctx is cancelled.
// It always closes jobs, so workers' range loops can end.
func producer(ctx context.Context, jobs chan<- int) {
	defer close(jobs)

	for j := 1; j <= numJobs; j++ {
		select {
		case <-ctx.Done():
			fmt.Println("Producer stopped:", ctx.Err())
			return

		case jobs <- j:
			fmt.Println("Produced Job:", j)
		}
	}

	fmt.Println("All jobs submitted")
}

// run starts the pool and blocks until ctx is cancelled and shutdown completes.
// Returns an error if the shutdown deadline was exceeded.
func run(ctx context.Context, timeout time.Duration) error {
	jobs := make(chan int, queueSize)

	// workCtx lets us cancel in-flight jobs if graceful shutdown takes too long.
	// It is NOT derived from ctx: Ctrl+C should stop the producer, not the workers.
	workCtx, cancelWork := context.WithCancel(context.Background())
	defer cancelWork()

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(workCtx, i, jobs, &wg)
	}

	producerDone := make(chan struct{})
	go func() {
		defer close(producerDone)
		producer(ctx, jobs)
	}()

	// Wait for Ctrl+C or SIGTERM.
	<-ctx.Done()
	fmt.Println("\nShutdown signal received!")

	// Everything is finished when the producer has closed jobs
	// and every worker has called wg.Done().
	allDone := make(chan struct{})
	go func() {
		<-producerDone
		wg.Wait()
		close(allDone)
	}()

	// Give workers a deadline to drain the queue.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-allDone:
		return nil

	case <-shutdownCtx.Done():
		fmt.Println("Shutdown deadline exceeded, cancelling in-flight jobs")
		cancelWork()
		<-allDone // workers return quickly once workCtx is cancelled
		return shutdownCtx.Err()
	}
}

func main() {
	// ctx is cancelled on the first Ctrl+C or SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// After the first signal, restore default behaviour so a second
	// Ctrl+C kills the process immediately.
	go func() {
		<-ctx.Done()
		stop()
	}()

	if err := run(ctx, shutdownTimeout); err != nil {
		fmt.Println("Server stopped with error:", err)
		os.Exit(1)
	}

	fmt.Println("Server stopped gracefully!")
}

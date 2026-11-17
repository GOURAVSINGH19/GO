package main

import (
	"fmt"
	"sync"
)

// RWMutex
// Now imagine your backend application has 100 users reading product information and only 2 users updating product information.
// Using a normal Mutex:
// - Only one reader can access the protected section at a time.
// - Other readers must wait.

// But we can improve this using RWMutex.
// RWMutex means Read-Write Mutex.

func main() {
	balance := 1000

	var wg sync.WaitGroup
	var rmu sync.RWMutex
	var mu sync.Mutex

	wg.Add(2)

	// withdraw := func(amount int) {
	// 	defer wg.Done()

	// 	// Locks the shared resource.
	// 	// If another goroutine already holds the lock, this goroutine waits.

	// 	mu.Lock()
	// 	balance = balance - amount

	// 	rmu.RLock()
	// 	fmt.Println("Reader 1", balance)
	// 	defer rmu.RUnlock()
	// 	defer mu.Unlock()
	// }

	// go withdraw(500)
	// go withdraw(500)

	go func() {
		defer wg.Done()

		rmu.RLock()
		fmt.Println("Reader 1:", balance)
		rmu.RUnlock()
	}()

	go func() {
		defer wg.Done()

		rmu.RLock()
		fmt.Println("Reader 2:", balance)
		defer rmu.RUnlock()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		balance = 2000
		fmt.Println("Balance updated:", balance)
		defer mu.Unlock()
	}()

	wg.Wait()

	fmt.Println("Final balance:", balance)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var cond *sync.Cond = sync.NewCond(&mtx)

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("Waiting for signal")
			cond.L.Lock()
			cond.Wait() // unlocks the condition variable's lock and suspends goroutine execution
			fmt.Printf("Worker %d woke up\n", i)
			cond.L.Unlock()
			fmt.Println("Done")
		}()
	}

	fmt.Println("Broadcasting after 2 seconds")
	time.Sleep(time.Second * 2)
	cond.Broadcast()

	wg.Wait()
	fmt.Println("ALL DONE")
}

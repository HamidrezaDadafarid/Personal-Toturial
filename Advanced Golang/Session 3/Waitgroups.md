```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// wg.Add(5) we can do this instead of incrementing it one by one. we have to know number of go routines.

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("Done with index", i) // the index will be random.
		}()
	}

	wg.Wait()
	fmt.Println("Everything is done")
}
```
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var num int
	var mtx sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock() // defer says what should be done just before leaving the function (e.g. return, panic).
			defer fmt.Println("Leaving...") // if we have multiple defers, they will be done in LIFO order.
			num++
			fmt.Println(num) // the prints, prints the num value in order from 1 to 100 because the print is locked too.
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("Result is", num)
}
```
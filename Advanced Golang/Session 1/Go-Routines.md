```go
package main

import (
	"fmt"
	"time"
)

func f(d time.Duration, message string) {
	time.Sleep(d)

	fmt.Println(message)
}

func main() {
	/* 
    when their sleep duration is the same, we don't know which one will be printed first.
    its random and based on scehduler. but here we are sure that f2 will be printed first because its sleep duration is much less than f1.
    */
	go f(time.Second*2, "Hello from f1")
	go f(time.Second*2, "Hello from f2")

	// we put this sleep to make sure all goroutines execute completely. if the main goroutine finishes, all goroutines will be killed.
    time.Sleep(time.Second * 4)
}

```
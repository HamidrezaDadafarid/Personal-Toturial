package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var num int64

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&num, 1) // all of the adds will be done sequentially without rce condition.
			fmt.Println(num)         // the print is not guarantied to be printed sequentially.
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("Result is", num)
}

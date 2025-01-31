package main

import (
	"fmt"
	"time"
)

func main() {
	var num int

	for i := 0; i < 100; i++ {
		go func() {
			num++ // if two go routines run this line at the same time, we will have two writes at the same time so the result is <= 100.
		}()
	}

	time.Sleep(time.Second * 2)
	fmt.Println("Result is", num)
}

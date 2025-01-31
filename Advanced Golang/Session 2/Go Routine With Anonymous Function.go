package main

import (
	"fmt"
	"time"
)

func main() {
	go func(number int) {
		time.Sleep(time.Second)
		fmt.Println(number)
	}(69)
	fmt.Println("After launching go routine")

	time.Sleep(time.Second * 2)
}

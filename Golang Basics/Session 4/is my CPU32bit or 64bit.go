package main

import "fmt"

func checkWordSize(desired uint) bool {
	return ^uint(0)>>(desired-1) == 1 // we use right shift because it is safe (unless left shift)
}

func main() {
	fmt.Println(checkWordSize(64))
}

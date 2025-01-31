package main

import "fmt"

func main() {
	// make is used for initializing slices, maps and channels

	var v1 []int
	fmt.Println(v1 == nil) // true

	v1 = append(v1, 10, 20, 30)
	fmt.Println(v1 == nil) // false

	v2 := make([]int, 10)         // a slice with length of 10 and capacity of 10
	fmt.Println(len(v2), cap(v2)) // 10, 10

	v2[9] = 99
	// v2[10] = 100        // panic. we should append it
	// fmt.Println(v2[10]) // panic again!

	v2 = append(v2, 100)
	fmt.Println(v2[10]) // no more panic!

	v3 := make([]int, 10, 20) // a slice with length of 10 and capacity of 20
	fmt.Println(v3[9])
	// fmt.Println(v3[10]) // slices range is from 0 to (length - 1)

	// when we append one or more elemtents to slice, if there is capacity for them it will be added. if not, a copy of that array with the length 2 times the current length will be made.
	for i := range 10 {
		v3[i] = i
	}
	fmt.Println(v3, len(v3), cap(v3))

	for i := range 10 {
		v3 = append(v3, i+10)
	}
	fmt.Println(v3, len(v3), cap(v3))

	v3 = append(v3, 20)
	fmt.Println(v3, len(v3), cap(v3))

}

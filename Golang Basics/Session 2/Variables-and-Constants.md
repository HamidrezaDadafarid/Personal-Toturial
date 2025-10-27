```go
package main

// there is a convention that variables of functions within the package that we don't want to export them, are named cammelCases.
// constants' are named UPPERCASE
// variables or functions that we want to export are named PascalCase

// global variable
var global int

func main() {
	// Defining variables
	var v1 int = 10
	var v2 string
	v3 := 34   // int by default
	v4 := 34.0 // float64 by default

	v5, v6, v7 := 1, "folan", false

	// variable grouping
	var (
		name   string
		family string
		age    uint
	)

	// constants: they are stored in code segment in memory so accessing them is so faster than when they are stored in stack or heap
	// The value of a constant must be assigned when you declare it.

	// Typed constant
	const C1 int = 100

	// Untyped constant
	const C2 = "Untyped Constant"

	// constant grouping
	const (
		PI = 3.14
		G  = 9.8
	)

	// Defining arrays
	var arr1 [5]int // An array of 5 integers, initialized to zero

	var arr2 [3]int = [3]int{1, 2, 3} // Array with predefined values
	var arr3 = [3]int{1, 2, 3}        // Array with predefined values
	arr4 := [3]int{1, 2, 3}           // Array with predefined values

	var arr5 = [...]int{10, 20, 30, 40}        // Compiler determines the size (4)
	var arr6 [4]int = [...]int{10, 20, 30, 40} // Compiler determines the size (4)
	arr7 := [...]int{10, 20, 30, 40}           // Compiler determines the size (4)

	var arr8 []int = []int{1, 2, 3}
	var arr8 = []int{1, 2, 3}
	arr9 := []int{1, 2, 3}

	var arr10 [5]int = [5]int{1: 100, 3: 200} // Creates [0, 100, 0, 200, 0]
	var arr11 = [5]int{1: 100, 3: 200}        // Creates [0, 100, 0, 200, 0]
	arr12 := [5]int{1: 100, 3: 200}           // Creates [0, 100, 0, 200, 0]

}
```
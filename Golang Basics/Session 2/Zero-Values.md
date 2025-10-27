```go
package main

import "fmt"

func main() {
	// zero values

	// numeric variables (uint(s), int(s), byte, rune) --> 0
	var v1 int
	var v2 uint
	var v3 byte
	var v4 rune
	var v5 complex64 // real part = 0, imaginary part = 0
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)

	// uintptr --> invalid address (address zero)
	var v6 uintptr
	fmt.Println(v6)

	// pointers --> nil
	var v7 *uint
	var v8 *bool
	var v9 *complex64
	var v10 *string
	fmt.Println(v7)
	fmt.Println(v8)
	fmt.Println(v9)
	fmt.Println(v10)

	// string --> empty string
	var v11 string
	fmt.Println(v11)

	// boolean --> false
	var v12 bool
	fmt.Println(v12)

	// array --> all elements are zero value of array's type
	var v13 [5]int
	fmt.Println(v13)

	// slices, maps, interfaces, channels --> nil
	var v14 []int
	var v15 map[string]bool
	var v16 interface{}
	var v17 chan int
	fmt.Println(v14 == nil)
	fmt.Println(v15 == nil)
	fmt.Println(v16 == nil)
	fmt.Println(v17 == nil)
	/*
	If you have an uninitialized slice of a certain type, like []string, and you use the fmt package to print it, the output will be [].
	This is because the fmt package uses reflection to determine the type of the slice and what to print.
	When the slice is uninitialized, the type is still known, so the fmt package is able to print an empty slice of that type.
	However, the underlying value of the slice is still nil, which can be misleading.
	*/

	/*
	structs --> For a struct type, if you declare a variable without explicitly initializing it, it will be assigned the zero value for its type.
	If the struct contains fields of numeric types, they will be set to zero values, and if it contains reference types (like slices or maps),
	those will be set to their respective nil values.
	*/

	type Person struct {
		firstName string
		lastName  string
		age       uint
	}
	var p Person
	fmt.Println(p.age)

}
```
package main

import "fmt"

// defining structs
type Person struct {
	firstName string
	lastName  string
	age       uint
	a         x // Named struct embedding
	x           // Asnonymous struct embeddin
}

type x struct {
	y int
}

func main() {
	// using named fields
	p1 := Person{
		firstName: "Hamidreza",
		lastName:  "Dadafarid",
		age:       22,
		a:         x{100},
		x:         x{200},
	}
	// we can only access y field of x struct like this in named struct embedding:
	fmt.Println(p1.a.y)

	// initializing a struct when the fields are provided in the same order as they are defined in the struct
	p2 := Person{"Hamidreza", "Dadafarid", 22, x{10}, x{20}}

	// we can access y field of x struct in two ways in anonymous struct embedding:
	fmt.Println(p2.x.y)
	fmt.Println(p2.y)

	// This declares a Person variable (p3) without initializing it. By default, the fields are assigned their zero values.
	var p3 Person
	p3.firstName = "Hamidreza"
	p3.lastName = "Dadafarid"
	p3.age = 22

	// p4 is a pointer to a Person and will initially hold the zero values for its fields
	// p4 := new(Person) // p4 = &Person{}

	// A nil pointer to a Person struct. You would need to assign it a valid reference (either through new(Person) or by assigning it the address of an existing Person).
	// var p5 *Person
}

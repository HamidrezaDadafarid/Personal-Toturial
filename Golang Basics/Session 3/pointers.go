package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func changeName(p *person) {
	p.firstName = "Changed first name"
}

func main() {
	var v1 int
	p := &v1
	*p = 40
	fmt.Println(*p, v1)
	v1 = 100
	fmt.Println(*p, v1)

	var p1 *person // zero value (nil)

	// panic!!!
	fmt.Println(p1.firstName) // equal to (*p1).firstName

	// not panic
	p2 := new(person) // equal tp &person{}
	fmt.Println(p2.firstName)

	p3 := &person{
		firstName: "Hamidreza",
		lastName:  "Dadafarid",
	}

	fmt.Println(p3.firstName)

	changeName(p3) // it is not pass by reference. we are just copying the address of p3 which is a number.
	fmt.Println(p3.firstName)
}

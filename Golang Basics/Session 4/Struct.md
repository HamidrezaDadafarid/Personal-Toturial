```go
package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       uint
}

// normal function
func fullName(p person) string {
	return p.firstName + " " + p.lastName
}

// method
/*
	methods will be converted to functions by go compiler
	we cannot have more than one receiver in a method.
	we cannot define methods on primitive types. just on user-defined types.
*/
func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

// if we don't use a pointer receiver here, we cannot see the set age outside of the method.
func (p *person) setAge(age int) {
	if age < 0 {
		age = 7
	}

	p.age = uint(age)
}

/*
This is a value receiver. The method operates on a copy of the person struct.
To get this value, Go needs a valid, non-nil person struct.
*/
func (p person) sayHello1() string {
	return "Hello"
}

/*
This is a pointer receiver. The method operates on a pointer to a person struct.
The pointer itself is just a memory address (or nil).
*/
func (p *person) sayHello2() string {
	return "Hello"
}

func main() {
	p1 := &person{
		firstName: "Hamidreza",
		lastName:  "Dadafarid",
	}

	fmt.Println(p1.fullName())
	p1.setAge(22)
	fmt.Println(p1.age)

	p2 := person{
		firstName: "Kiana",
		lastName:  "Sherafati",
	}

	fmt.Println(p2.fullName())
	// Syntactic Sugar: if we don't call the method on a pointer to that struct, the struct's address will be passed to the method.
	p2.setAge(24)
	fmt.Println(p2.age)

	var p3 *person

	/*
	it will panic because we are trying to derference the nil pointer then pass it to the method.
	*/
	fmt.Println(p3.sayHello1())

	/*
	it will work ok because we don't need to dereference the pointer.
	This call is safe as long as the method doesn't try to use any of the struct's fields.
	sayHello2 method just returns "Hello"; it never tries to access.
	*/
	fmt.Println(p3.sayHello2())
}

```
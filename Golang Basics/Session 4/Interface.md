```go
package main

type Human interface {
	Name() string
	Age() uint
	SayHello() string
}

type Person struct {
	firstName string
	lastName  string
	age       uint
}

// Person struct has implemented all Human interface methods. so it is a Human by default.
func (p Person) Name() string {
	return ""
}

func (p Person) Age() uint {
	return 0
}

func (p Person) SayHello() string {
	return ""
}

func main() {

	var h1 Human
	var h2 Human

	// if the person does not implement one of the interface methods, we can not assign a person or a pointer to a person to a human.
	// if at least one of the methods have pointer receiver, we can only assign a pointer to a person to a human not the value of it.
	h1 = Person{"Hamidreza", "Dadafarid", 22}
	h2 = &Person{"Hamidreza", "Dadafarid", 22}

}

```
package main

import "fmt"

func main() {
	v1 := map[string]string{
		"firstName": "Hamidreza",
		"lastName":  "Dadafarid",
	}

	if _, exists := v1["firstName"]; exists {
		fmt.Println("first name key exists!")
	} else {
		fmt.Println("first name key does not exists")
	}

	// we use delete buit-in function to delete a key from our map.
	delete(v1, "firstName")

	if _, exists := v1["firstName"]; exists {
		fmt.Println("first name key exists!")
	} else {
		fmt.Println("first name key does not exists")
	}
}

package main

import "fmt"

func main() {
	// keys in maps must be comparable, there is no limit in values' type

	var v1 map[string]int // nil
	fmt.Println(v1 == nil)

	v2 := make(map[string]bool) // an empty map with len 0
	fmt.Println(len(v2))

	v3 := map[string]bool{
		"Hamidreza": true,
		"Ali":       false,
		// .
		// .
		// .
	}

	fmt.Println(v3["Hamidreza"])
	fmt.Println(v3["Folan"]) // if the key is not present in map, it returns zero value of its value type. we should do:

	value, exists := v3["Hamidreza"]

	if exists {
		fmt.Println("key exists! The value is", value)
	} else {
		fmt.Println("key does not exists!. The value is", value)
	}

	value, exists = v3["Folan"]

	if exists {
		fmt.Println("key exists! The value is", value)
	} else {
		fmt.Println("key does not exists!. The value is", value)
	}
}

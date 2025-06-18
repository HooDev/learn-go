package main

import "fmt"

func main() {
	// Maps
	ages := map[string]int{
		"Daniel": 38,
		"Mary":42,
	}
	fmt.Println(ages)

	fmt.Println("Mary is", ages["Mary"])

	ages["Bob"] = 70

	fmt.Println(ages)

	ages["Steve"] = 50

	// exists
	value, exists := ages["Steve"]

	if exists {
		fmt.Println("Steve is", value)
	} else {
		fmt.Println("We have no Steve")
	}

	delete(ages, "Bob")

	fmt.Println(ages)
}

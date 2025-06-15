package main

import "fmt"

func main() {
	sayHello("Daniel")

	added := addNumbers(2, 4)

	fmt.Println("our result", added)

	// Divide by Zero
	divided, errors := divideNumbers(10, 0)
	fmt.Println(divided, errors)

	// Working division
	divided, errors = divideNumbers(5.55, 8.88)
	fmt.Println(divided, errors) 

	divided, err := divideNumbers(23, 10)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(divided)
	}
}

func sayHello(name string) {
	fmt.Println("Hello!", name)
}

func addNumbers(a int, b int) int {
	return a + b
}

func divideNumbers(a, b float64) (float64, string) {
	if b == 0  {
		return 0, "Cannot divide by Zero"
	}

	return a / b, "" 
}

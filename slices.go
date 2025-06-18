package main

import "fmt"

func main() {
	// Slices
	names := []string{"Daniel", "Mary"}

	fmt.Println(names)

	// Append
	names = append(names, "Mike")

	fmt.Println(names)

	numbers := [8]int{0:1, 3: 808, 7:888}

	fmt.Println(numbers)

	sliceFromArray := numbers[2:5]

	fmt.Println(sliceFromArray)
}


package main

import "fmt"

func main() {
	
	// Conditionals
	// If, Else If, Else
	age := 1
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You're a teenager")
	} else {
		fmt.Println("Sorry, you're not old enough, yet")
	}

	// Logical Operators
	score := 85

	if score >= 70 && score <= 100 {
		fmt.Println("Congrats on the good score!")
	}

	// Logical Operators: && (and), || (or), ! (not)
	if score > 100 || score < 0 {
		fmt.Println("Score Invalid")
	}

	// Switch Statement (score)
	switch {
		case score >= 90:
			fmt.Println("High Distinction")
		case score >= 80:
			fmt.Println("Distinction")
		case score >= 70:
			fmt.Println("Credit")
		case score >= 50:
			fmt.Println("Pass")
		default:
			fmt.Println("Fail")
	}

	// Switch Statement
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Thank God it's Friday!")
	default:
		fmt.Println("Another day!")
	}
}

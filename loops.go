package main

import "fmt"

func main() {
	fmt.Println("1. Basic For Loop")
	for i := 0; i < 10; i++ {
		fmt.Println("i is:", i)
	} 


	fmt.Println("2. While Style For Loop")
	i := 0
	for i < 20 {
		fmt.Println("i is:", i)
		i++
	}

	fmt.Println("4. Break and Continue")
	for i := 0; i < 20; i++ {
		// continue
		if i % 2 == 0 {
			continue
		}
		// break
		if i == 9 {
			break
		}
		fmt.Println("i is", i)
	}  
	
	fmt.Println("3. Infinite Loop")
	for {
		fmt.Println("Don't stop me now!")
		break
	}

	// Range
	fmt.Println("5. Looping with Range")
	names := []string{"Steve", "Mary", "Robert"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, name := range names {
		fmt.Println(name)
	}

}


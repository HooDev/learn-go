package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	age := 30
	fmt.Println("address of age:", &age)
	fmt.Println("age variable, before function", age)

	changeAge(age)

	fmt.Println("age variable, after reg function", age)
	
	changeAgePtr(&age)

	fmt.Println("age variable, after ptr function", age)

	daniel := Person{Name: "Daniel", Age: 30}
	fmt.Println(daniel)

	daniel.HaveBirthday()
	fmt.Println(daniel)
}

func changeAge(a int) {
	a = 35
}

func changeAgePtr(a *int) {
	*a = 35
}

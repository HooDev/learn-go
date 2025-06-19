package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	daniel := Person{Name: "Daniel", Age: 31}

	fmt.Println(daniel)

	var alice Person
	fmt.Println(alice)

	alice.Name = "Alice"
	alice.Age = 40

	fmt.Println(alice)

	fmt.Println("Alice is", alice.Age)

	alice.Greet()
	alice.HaveBirthday()

	fmt.Println("Alice is", alice.Age)
}

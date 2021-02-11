package main

import "fmt"

type human struct {
	name   string
	age    int
	weight float64
}

// value receiver method (immutable)

func (p human) intro() string {
	return "Hello my name is " + p.name
}

// pointer receiver method

func (p *human) changeName(name string) string {
	p.name = name
	return p.name
}

// Inheritance

type man struct {
	human
}

type woman struct {
	human
}

func main() {

	bob := human{name: "Bob", age: 28, weight: 62}

	fmt.Println("Bob ", bob)

	fmt.Println(bob.intro())

	bob.changeName("Alex")
	fmt.Println(bob.intro())

	alice := woman{human{name: "Alice", age: 29, weight: 57}}

	fmt.Println(alice.intro())
}

package main

import "fmt"

type human struct {
	name   string
	age    int
	weight float64
}

func (p human) intro() string {
	return "Hello my name is " + p.name
}

func main() {

	bob := human{name: "Bob", age: 28, weight: 62}

	fmt.Println("Bob ", bob)

	fmt.Println(bob.intro())

}

package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

func add(a, b int) int {

	return a + b
}

// Return mutiple values

func sqr(a int) (int, string) {

	d := a * 2
	s := fmt.Sprintf("square  is %d", d)

	return d, s

}

// Function clousure return another function

func intro(name string) func() {

	fmt.Println("Hi " + name)

	return func() {

		fmt.Println("Good bye ", name)

	}
}

func bio(name string, age int) (string, int) {

	return fmt.Sprintf("My name is %s", name), age

}

func printName(name *string) {

	*name = "Bean"

	fmt.Println("My name is ", *name)
}

// variable number of params

func sum(values ...int) int {

	var sum int

	for _, v := range values {
		sum = sum + v
	}
	return sum

}

// Named return

func getMyName() (name string) {

	name = "Sarath"

	return
}

func divide(a, b float64) (float64, error) {

	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide with zero")
	}

	return a / b, nil

}
func main() {

	sayHello()

	fmt.Println("Sum of 1 and 2 is", add(1, 2))

	v, text := sqr(4)

	fmt.Println(v, text)

	bye := intro("Sarath")

	bye()

	// mutiple return value
	bio, age := bio("Bob", 20)

	fmt.Println(bio, age)

	// Pointer params

	myname := "Teddy"

	printName(&myname)

	fmt.Println("my name variable changed : ", myname)

	fmt.Println(sum(1, 2, 3, 4, 4, 5, 5, 6))

	scores := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(scores...))

	fmt.Println("My name is ", getMyName())

	result, err := divide(2.0, 1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Division result ", result)

	// Anonymous functions

	func() {
		fmt.Println("This is a anonymous function ")
	}()

	// Anonymous functio as a variable

	f := func(name string) {
		fmt.Println("Hello ", name)
	}

	f("Bob")

	// Anonymous function with signature

	var goodbye func(name string) string

	goodbye = func(name string) string {

		fmt.Println("Good bye ", name)
		return "Good bye"
	}

	goodbye("Alex")

}

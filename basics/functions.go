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

// Variable number of arguments

func sumOfNumbers(nums ...int) int {

	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum

}

func main() {

	sayHello()

	fmt.Println("Sum of 1 and 2 is", add(1, 2))

	v, text := sqr(4)

	fmt.Println(v, text)

	total := sumOfNumbers(1, 2, 3)

	fmt.Println("Ttal sum of numbers are ", total)
}

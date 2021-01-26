package main

import "fmt"

func main() {

	var x int = 10

	y := &x

	fmt.Println("memory location of x ", y)

	// Access value of pointer

	fmt.Println("Print *y", *y)

	// Typed pointer variables
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	// Change value of a
	a = 27
	fmt.Println(a, b)

	// Change value o *b
	*b = 10
	fmt.Println(a, b)

	// maps and slices are pointer by default
}

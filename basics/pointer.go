package main

import "fmt"

func main() {

	var x int = 10

	y := &x

	fmt.Println("memory location of x ", y)

	// Access value of pointer

	fmt.Println("Print *y", *y)
}

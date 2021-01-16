package main

import (
	"fmt"
)

func main() {

	//Unlike an array type, a slice type has no specified length

	//A slice literal is declared just like an array literal, except you leave out the element count

	letters := []string{"a", "b", "c"}

	fmt.Println("Letters ", letters)

	// A slice can be created with the built-in function called make, which has the signature
	s1 := make([]string, 3)

	fmt.Println("Slice s1 ", s1)
	var s2 = []int{1, 2, 3}
	fmt.Printf("value %v and type %T", s2, s2)
	fmt.Println("\n Lenght :", len(s1))
	fmt.Println("\n Capacity :", cap(s1))

	// Slice operations
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //Slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  //Slice first 6 elements
	e := a[3:6] // Slice the 4,5,and 6th element

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("Capacity of e ", cap(e))

	// Slice operations
	// Add (append )

	x := []int{}
	fmt.Println("X Length", len(x), "Capacity", cap(x))
	x = append(x, 1)
	fmt.Println("X Length", len(x), "Capacity", cap(x))

	x = append(x, 2, 3, 4, 5)
	fmt.Println("X Length", len(x), "Capacity", cap(x))

	// Pop

	y := x[1:]
	fmt.Println("Remove first element ", y)

	z := x[:len(x)-1]
	fmt.Print("Removed last element ", z)

	middle := append(x[:2], x[3:]...)

	fmt.Println("Remove from middle ", middle)
	fmt.Println("X ", x)

}

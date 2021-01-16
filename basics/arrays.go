package main

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 5

	fmt.Println(a)

	// initialize array with values

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b ", b)

	var c = [6]int{1, 2, 3}

	fmt.Println("Array c ", c)

	// Get array length

	fmt.Println("Array length ", len(c))

	// Grades
	grades := [...]int{97, 85, 93} // Size from initial data
	fmt.Printf("Grades %v", grades)

	// String arrays
	var students [3]string
	fmt.Println("\nSudents ", students)
	students[0] = "Lisa"
	fmt.Println("\nSudents ", students)
	students[2] = "Alice"
	students[1] = "Bob"
	fmt.Println("\nSudents ", students)

	//Multi Dimention Array

	// Array operations

	var players = [3]int{1, 2, 3}
	ref := &players
	fmt.Println(*ref)

}

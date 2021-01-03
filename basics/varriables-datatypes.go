package main

import "fmt"

func main() {
	var i int = 10
	var name string = "Sarath"
	var points = 10.50
	var isEnabled = true

	fmt.Printf("Number is %v %T", i, i)
	fmt.Printf("\nName is %v %T", name, name)
	fmt.Printf("\nPoints is %v %T", points, points)
	fmt.Printf("\nIs Enabled %v %T", isEnabled, isEnabled)

	// Shortcut variable declaration

	theURL := "https://golang.org"
	fmt.Printf("\n" + theURL)

	// Numbers;
	var count int32 = 100
	var pi float32 = 3.14

	fmt.Printf("\nCount %v", count)

	fmt.Printf("\nValue of pi %v", pi)

}

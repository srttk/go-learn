package main

import "fmt"

func main() {

	/* Boolean (bool)
	Represents a truth value. It can be either `true` or `false`.
	*/

	var isValid bool = true
	fmt.Printf("Is valid %v \n", isValid)
	// Numeric data types

	/* Integers
	integers can be signed (`int`) or unsigned (`uint`). The size of these types depends on the architecture (32 or 64 bits). Go also provides specific sizes like `int8`, `int16`, `int32`, and `int64`.
	*/

	var year int = 2024
	fmt.Printf("Year is %v \n", year)

	/* Float
	float32, float64 , Used for numbers with decimal points. The default type for floating-point numbers is `float64`.
	*/

	var price float64 = 54.50

	fmt.Printf("Price is %v \n", price)

	/* String
	A sequence of characters. Strings are immutable in Go.
	*/

	var country string = "India"
	fmt.Printf("Country %v\n", country)

	/* Complex numbers
	complex64, complex128 : Used for complex numbers with real imaginary parts.

	*/

	var c complex128 = 5173 + 2i

	fmt.Printf("Sample  compex number %v", c)

}

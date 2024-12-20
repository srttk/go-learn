package main

import "fmt"

func main() {



	/* Variables

	In Go, variables are declared using the `var` keyword followed by the variable name, type, and an optional initializer. For example:
	*/


	var x int
	var y int = 10
	var z = 20

	fmt.Printf("Variable values : x=%v, y=%v, z=%v \n", x,y,z)


	/*
	Multiple variables can be declared in one line:
	*/

	var a,b,c int = 1,2,3

	fmt.Printf("Values of a=%v, b=%v, c=%v \n", a,b,c)


	/* Shorthand syntax
	
	When you use the shorthand declaration, Go automatically determines the type of the variable based on the value assigned. However, once a type is assigned, it cannot be changed.



	*/

	shorthand:="This is a shorthand variable"

	fmt.Printf("%v \n",shorthand)

	/* shorthand multiple variable*/
	name,place:="Luke", "Tatoonie"

	fmt.Printf("%v from %v \n", name, place)


	/*shorthand - scope limitations 
	
	The short declaration operator can only be used inside functions; it cannot be used at the package level. For package-level declarations, you must use the `var` keyword. Additionally, at least one of the variables in a multi-variable declaration must be new for the shorthand to be valid.
	*/




}
package main

import "fmt"

func main() {

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 7%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if n := 3; n < 0 {
		fmt.Println("Number is negative")
	} else if n < 10 {
		fmt.Println("Single digit number")
	} else {
		fmt.Println("Multi digit number")
	}

}

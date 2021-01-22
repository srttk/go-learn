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

	// map colors
	colors := map[string]string{"red": "#ff0000", "green": "#00ff00"}

	// Initializer syntax
	if _, ok := colors["green"]; ok {

		fmt.Println("Color green exists, the code is :", colors["green"])
	}

	// Number guessing game
	number := 50
	guess := 67

	if guess < 0 || guess > 100 {
		fmt.Println("The guess must between  1 and 100")

	} else {

		if number > guess {
			fmt.Println("Too low")
		} else if number < guess {
			fmt.Println("Too large")
		} else {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(number >= guess, number <= guess, number != guess)
}

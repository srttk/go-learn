package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch ")
	i := 4
	// Baisc
	switch i {
	case 1:
		fmt.Println("o-n-e")
	case 2:
		fmt.Println("t-w-o")
	case 3:
		fmt.Println("t-h-r-e-e")
	}

	fmt.Println("Time ", time.Now())
	fmt.Println("Weekday ", time.Now().Weekday())

	// With default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend 😇")
	default:
		fmt.Println("It's working day 🙁")
	}

	//
	t := time.Now()
	fmt.Println("Time now ", t)

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning 🕒")
	default:
		fmt.Println("Afternoon 🕓")
	}

	// multiple case

	switch 2 {
	case 2, 4, 6:
		fmt.Println("two four or six")
	case 1, 3, 5:
		fmt.Println("one three or five")

	default:
		fmt.Println("another number")
	}

	// inittializer example
	val := 6
	switch result := val * 10; result {
	case 20, 30, 40:
		fmt.Println("20,30,40")
	case 50, 60, 70:
		fmt.Println("50,60,70")
	default:
		fmt.Println("Other result")
	}

	// Tagless and fallthrough

	j := 10

	switch {
	case j <= 10:
		fmt.Println("j less than or equal to 10")
		fallthrough
	case j <= 20:
		fmt.Println("j less than or equal to 20")

	default:
		fmt.Println("j greatr than 20")
	}

}

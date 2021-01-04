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
}

package main

import (
	"fmt"
	"log"
)

/**
panic like throw error

*/

func panicer() {
	fmt.Println("Function start")
	defer func() {

		if err := recover(); err != nil {
			log.Println("Error ", err)
		}

	}()
	panic("Something bad happend")
	fmt.Println("Function end")
}

func main() {
	fmt.Println("App Started")

	// result := 1 / 0 -> returns a panic

	panicer()

	fmt.Println("App ends")
}

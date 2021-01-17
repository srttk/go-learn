package main

import (
	"fmt"
)

func main() {

	var names = [5]string{"alex", "stuart", "bob", "gru", "rob"}

	for index, name := range names {
		fmt.Println(" name ", name, " , index :", index)
	}

	// Sum of an slice

	numbers := []int{9, 8, 7, 8, 8}
	var sum int

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum of numbers : ", sum)

	// iterate maps
	extensions := map[string]string{"js": "javascript", "ts": "typescipt"}

	for key, value := range extensions {

		fmt.Println(value, ".", key)
	}

	// Strings
	for i, v := range "sarath" {
		fmt.Println(i, " -> ", v)
	}
}

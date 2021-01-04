package main

import "fmt"

func main() {
	var i int = 1

	for i <= 3 {
		fmt.Println("Value is ", i)
		i = i + 1
	}

	for j := 5; j <= 9; j++ {
		fmt.Println("Value of j ", j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := 0; n <= 9; n++ {

		if n%2 == 0 {
			continue
		}

		fmt.Println(n, "is odd")
	}

}

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10000000000; i++ {
		//
	}

	fmt.Println(time.Since(start))
}

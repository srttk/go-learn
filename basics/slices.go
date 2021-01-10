package main

import (
	"fmt"
)

func main() {

	s1 := make([]string, 3)

	fmt.Println("Slice s1 ", s1)

	var s2 = []int{1, 2, 3}
	fmt.Printf("value %v and type %T", s2, s2)

}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Printf("Threads %v \n ", runtime.GOMAXPROCS(-1))

	// Example 1
	go sayHello()
	time.Sleep(100 * time.Microsecond)

	// Example 2

	var msg string = "World"

	go func() {
		fmt.Println(msg)
	}()

	msg = "Universe"

	fmt.Println(msg)

	time.Sleep(100 * time.Microsecond)

	// Example 3

	//TODO: More examples with sync.waitGroup

}

func sayHello() {
	fmt.Println("Hello")
}

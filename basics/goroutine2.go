package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		count("Sheep")
		wg.Done()
	}()

	wg.Wait()

}

func count(label string) {

	for i := 1; i < 10; i++ {
		fmt.Println(label, i)
	}
}

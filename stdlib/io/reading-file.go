package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Method - 1

	data, error := ioutil.ReadFile("./stdlib/io/files/data.txt")
	check(error)

	fmt.Println(string(data))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

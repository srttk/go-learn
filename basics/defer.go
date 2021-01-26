package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Defer helps to Closing opened resourced after process
func exampleClosingResources() {

	res, err := http.Get("http://www.google.com/robots.txt")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)

	}

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(robots))

}

func main() {

	exampleClosingResources()

	fmt.Println("First")
	defer fmt.Println("Middle")
	fmt.Println("Last")

	text := "Hello"
	defer fmt.Println(text)
	text = "World"

}

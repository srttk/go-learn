package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go updateNews("Boom", ch)

	news := <-ch

	fmt.Println("Hot News : ", news)

	fmt.Scanln()

}

func updateNews(content string, c chan string) {

	time.Sleep(4 * time.Second)
	c <- content

}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Http : get request")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln("Error")
	}

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatalln("Request body read error")
	}

	stringBody := string(body)

	fmt.Println(stringBody)
}

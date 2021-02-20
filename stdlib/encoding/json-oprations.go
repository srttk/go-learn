package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Author      Author `json:"author"`
	Year        int    `json:"year"`
}

type Author struct {
	Name string `json:"name"`
}

func main() {

	// Struct -> Json string (json.Marshal or json.MarshalIndent)
	author := Author{Name: "Yuval Noah Harari"}
	book := Book{Title: "Sapiens", Description: "A Brief History of Humankind", Author: author}

	fmt.Printf("Book : %+v", book)

	data, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(data))

	// json string -> struct (json.Unmasrshal)
	jsonString := `{"title":"Maxim Gorky",
	"desc":"Mother is a novel written by Maxim Gorky in 1906 about revolutionary factory workers","year":1906,
	"author": {"name":"Maxim Gorky"}}`

	var book2 Book

	error := json.Unmarshal([]byte(jsonString), &book2)

	if error != nil {
		fmt.Println("Conversion error", error)
		return
	}

	fmt.Printf("unmarshed data %+v \n", book2)

	fmt.Println("Book 2 name : ", book2.Title)
}

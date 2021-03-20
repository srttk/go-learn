package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Provide a directry")
	}

	files, error := ioutil.ReadDir(args[0])

	if error != nil {
		fmt.Println("Directory read error")
		return
	}

	// fmt.Println(files)

	var emptyFiles []string

	for _, file := range files {

		name := file.Name()
		if file.Size() == 0 {

			emptyFiles = append(emptyFiles, name)
		}
	}

	if len(emptyFiles) == 0 {
		fmt.Println("✅  No empty files found ")
		return
	}

	fmt.Printf("⚠️  %d empty files found \n", len(emptyFiles))

	for _, filename := range emptyFiles {

		fmt.Println(" 	► ", filename)
	}

}

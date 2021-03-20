package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	args := os.Args[1:]

	// Set default directory as current directory
	directory := "."
	directoryLabel := "current"

	if len(args) != 0 {
		directory = args[0]
	}

	switch directory {
	case ".":
		directoryLabel = "current"
	case "..", "../":
		directoryLabel = "parent"
	default:
		directoryLabel = directory
	}

	files, error := ioutil.ReadDir(directory)

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
		fmt.Println("✅  No empty files found in ", directoryLabel, " directory")
		return
	}

	fmt.Printf("⚠️  %d empty files found in %s directory \n", len(emptyFiles), directoryLabel)

	for _, filename := range emptyFiles {

		fmt.Println(" 	► ", filename)
	}

}

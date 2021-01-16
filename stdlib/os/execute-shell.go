package main

import (
	"fmt"
	"os/exec"
)

func main() {

	gopath, err := exec.LookPath("go")

	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println("Result ", gopath)
}

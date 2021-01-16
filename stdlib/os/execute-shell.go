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

	// Execute

	cmd := exec.Command("ls", "-l")
	stdout, err := cmd.Output()

	fmt.Println(string(stdout))

}

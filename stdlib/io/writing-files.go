package main

import "io/ioutil"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	data := []byte("Hello World")

	error := ioutil.WriteFile("./stdlib/io/files/data2.txt", data, 0644)

	checkError(error)

}

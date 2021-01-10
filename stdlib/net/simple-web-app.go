package main

import (
	"fmt"
	"net/http"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm Sarath")
}

func getJSONAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, "{\"message\":\"hello\"}")
}

func main() {

	fmt.Println("Simple web app")

	http.HandleFunc("/", getIndex)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/json", getJSONAPI)

	http.ListenAndServe(":8000", nil)

}

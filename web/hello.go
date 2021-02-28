package main

import (
	"fmt"
	"net/http"
)

// PORT Number
const PORT = ":80"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Get query params
		msg := r.URL.Query().Get(("msg"))

		// Set Headers
		w.Header().Set("Content-Type", "text/html")

		// Send back response
		fmt.Fprintf(w, "Hello World %s  and message is %s", r.URL.Path, msg)
	})
	fmt.Println("Server listening at port " + PORT)
	error := http.ListenAndServe(PORT, nil)

	if error != nil {
		fmt.Println("Error : ", error)
	}

}

package main

import (
	"fmt"
)

func main() {

	var popularity = map[string]int{
		"react":   5,
		"angular": 2,
		"vue":     3,
		"svelte":  4,
	}
	fmt.Println("Maps", popularity)
	// Add another item
	popularity["angularjs"] = 1
	fmt.Println("Maps", popularity)

	// Accessing item

	fmt.Println("Itm : ", popularity["react"])

	// Aceessing non existing item

	value, ok := popularity["express"]

	fmt.Println("value ", value, ok)

	langs := make(map[string]int)

	langs["node"] = 5
	langs["php"] = 4

	// Length

	fmt.Println("Length of map langs ", len(langs))

	// Remove from map

	delete(langs, "php")

	fmt.Println("after delete ", langs)

}

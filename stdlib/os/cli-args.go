package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Map of available emojies
var emojies = map[string]string{
	"fire":  "🔥",
	"smile": "😀",
	"poo":   "💩",
	"wave":  "👋",
}

func printEmoji(emojiName string, count int) {

	emoji, isExist := emojies[emojiName]

	if isExist == false {
		fmt.Printf("👎🏻 No Emoji found named %v 🚫", emojiName)

	}

	for i := 0; i < count; i++ {
		fmt.Print(emoji)
	}
	println("")
}

func main() {

	args := os.Args
	//fmt.Println(args)
	if len(args) == 1 {
		fmt.Printf("Ooops")
		return
	}

	command := args[1]

	count := 1

	command = strings.ToLower(command)

	if len(args) > 2 {

		num, error := strconv.Atoi(args[2])
		if error != nil {
			fmt.Println("Invalid count", error)
		}

		count = num
	}

	printEmoji(command, count)
}

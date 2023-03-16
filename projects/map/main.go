package main

import "fmt"

func main() {

	// colors := make(map[string]string)

	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"

	// delete(colors, "white")

	fmt.Println(colors)
}

func printMap (c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
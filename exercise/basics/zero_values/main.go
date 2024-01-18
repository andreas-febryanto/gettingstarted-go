package main

import "fmt"

func main() {
	// Zero values each type
	// numeric: 0
	// boolean: false
	// string: ""
	// pointer: nil

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
package main

import "strconv"

func main() {
	// if, else if and else
	price, inStock := 100, true

	if price > 80 {
		println("Too expensive")
	}

	if price <= 100 && inStock {
		println("Buy it!")
	}

	if price < 100 {
		println("It's cheap")
	} else if price == 100 {
		println("On the edge")
	} else {
		println("It's expensive")
	}

	// simple if
	if i, err := strconv.Atoi("10"); err == nil {
		println(i)
	} else {
		println(err)
	}


}
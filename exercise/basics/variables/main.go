package main

import (
	"fmt"
)

func main() {
	var age int = 30
	fmt.Println("Age:", age)
	
	// type inference
	var name = "Dan"
	fmt.Println("Name:", name)

	// short hand
	s := "Learning golang!"
	fmt.Println(s)

	// multiple variables
	car, cost := "audi", 50000
	fmt.Println(car, cost)

	// at least one need to be new
	car, year := "bmw", 2018
	_ = year // ignore a variable

	// another way to create multiple variables
	var (
		salary float64
		firstName string
		gender bool
	)
	fmt.Println(salary, firstName, gender)

	// multiple assignments
	var i, j int
	i, j = 5, 8
	j, i = i, j // swap i and j
	fmt.Println(i, j)

	// multiple assignments expression
	sum := 5 + 2.3
	fmt.Println(sum)

	// Constants rules
	// 1. You can't change a constant
	// 2. You can't initialize a constant at runtime
	// const power = math.Pow(2, 3)
	// 3. You can't initialize a constant using a variable
	// t := 5
	// const tc = t
	
	// const must be given a value
	const days int = 7
	const n, m int = 4, 5
	const n1, m1 = 6, 7

	// const block, initialize multiple const
	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)
	fmt.Println(min1, min2, min3)
}
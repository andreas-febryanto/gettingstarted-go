package main

import "fmt"

func main() {
	// Numeric types
	// int, int8, int16, int32, int64 => signed integers
	// uint, uint8, uint16, uint32, uint64 => unsigned integers
	// float32, float64 => floating point numbers
	// complex64, complex128 => complex numbers
	// byte => alias for uint8
	// rune => alias for int32
	// bool => true or false
	// string => string of characters

	// Composite types
	// array => fixed length list of values. Must be of same type
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", numbers)

	// slice => dynamic length list of values
	var cities = []string{"London", "New York", "Tokyo"}
	fmt.Printf("%T\n", cities)

	// map => unordered list of key-value pairs
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances)

	// struct => collection of fields
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	// pointer => pointer to another value in the program
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr)

	// function => function with optional name
	fmt.Printf("%T\n", main)

	// interface => set of methods

	// arithmatic, assignment, comparison, logical, bitwise operators
	// arithmetic operators: +, -, *, /, %, ++, --
	// assignment operators: =, +=, -=, *=, /=, %=
	// comparison operators: ==, !=, <, <=, >, >=
	// logical operators: &&, ||, !
	// bitwise operators: &, |, ^, <<, >>
	// operator for pointers: &, *
	// operator for channels: <-

	// Defined types
	type age int // int is its underlying type
	type veryOldAge age 

}
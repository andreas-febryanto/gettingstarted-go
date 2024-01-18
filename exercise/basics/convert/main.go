package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1
	
	x = x * int(y)
	fmt.Println(x)

	x = int(float64(x) * y)
	fmt.Println(x)

	var a = 5 
	var b int64 = 2
	a = int(b)
	fmt.Println(a)

	s := string(99)
	fmt.Println(s)

	// convert float to string
	// s1 := string(44.2)
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	// convert string to float
	var f1, err = strconv.ParseFloat("44.2", 64)
	_ = err
	fmt.Println(f1)
	// string to int
	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s)


}
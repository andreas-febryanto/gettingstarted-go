package main

import "fmt"


func main() {
	// for i := 0; i < 5; i++ {
	// 	println(i)
	// }

	// for i := 0; i < 5; {
	// 	println(i)
	// 	i++
	// }

	// j := 5
	// for j >= 0 {
	// 	println(j)
	// 	j--
	// }

	// for i := 0; i <= 10; i++ {
  //   if i%2 != 0 {
  //       continue
  //   }
  //   println(i)
	// }

	// count := 0

	// for i:=0; true; i++ {
	// 	if i % 13 == 0 {
	// 		fmt.Printf("%d is divisible by 13\n", i)
	// 		count++
	// 	}

	// 	if count == 10 {
	// 		break
	// 	}
	// }

	// Labels are used to break out of nested loops
	people := [5]string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
	friends := [2]string{"Mark", "Dave"}

	outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	println("Next instruction after the break.")

	
}
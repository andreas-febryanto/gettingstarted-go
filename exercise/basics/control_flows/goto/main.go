package main

// Goto statement is not recommended to use in Go and any other programming languages
// It is not recommended because it makes code hard to read, understand, and code hard to maintain

func main() {
	i := 0

	loop:
		if i < 5 {
			println(i)
			i++
			goto loop
		}

		// Not allowed after goto statement, there is variable declaration
		// goto todo
		// x := 5
		// todo:
		// 	println("something here")
}
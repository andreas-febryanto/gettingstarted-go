package main

func main() {
	language := "golang"

	switch language {
	case "python":
		println("You are learning Python")
	case "Go","golang":
		println("You are learning Golang")
	default:
		println("Any other programming languages")
	}
}
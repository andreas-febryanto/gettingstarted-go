## Intro

Go is an open-source systems programming language initially developed as an internal Google project that went public back in 2009.

> Although the official name of the language is Go, it is sometimes (wrongly) referred to as Golang. The official reason for this is that go.org was not available for registration and golang.org was chosen instead. The practical reason for this is that when you are querying a search engine for Go-related information, the word Go is usually interpreted as a verb.

## Advantages & Disadvantages

- Go has no direct support for object-oriented programming, which is a popular programming paradigm

- Although goroutines are lightweight, they are not as powerful as OS threads.

- Go will not allow you to perform any memory management manually.

## Go doc & godoc utilities

Two of these tools are the go doc subcommand and godoc utility, which allow you to see the documentation of existing Go functions and packages without needing an internet connection. Or you can access here: [Go Packages](https://pkg.go.dev/)

> ex: go doc fmt

## Quick start

1. Install Go Runtime -> [Downloads - The Go Programming Language](https://golang.org/dl/)

2. Run go in terminal

3. go run <filename.go>

```go
package main

import ("fmt")

func main() {
 fmt.Println("Hello World!")
}
```

### Functions

Each Go function definition begins with the func keyword followed by its name, signature and implementation.

### packages

Go programs are organized in packages—even the smallest Go program should be delivered as a package. The package keyword helps you define the name of a new package.

When importing a package, Go checks the GOPATH and GOROOT environment variables. The GOPATH points to the Go workspace and it is defined during the installation Similarly, GOROOT points to a custom Go installation. This variable should not be required unless a custom installation is done. The Go compiler will first check the GOROOT and then the GOPATH when importing a package.

> Package == Project == Workspace
> 
> - Executable package
>   
>   package main(special) -> Defines a package that can be compiled and then *executed*. **Must have a func called 'main'**
> 
> - Reusable package
>   
>   package calculator -> Defines a package that can be used as a dependency(helper code)

###### Import from third party

The import keyword is used for importing other Go packages in your Go programs in order to use some or all of their functionality. A Go package can either be a part of the rich Standard Go library or come from an external source. Packages of the standard Go library are imported by name (os) without the need for a hostname and a path, whereas external packages are imported using their full internet paths, like github.com/spf13/cobra.

Go forces code transparency by only compiling source code. This means that in order to import third-party packages, the source code must be locally available. Before import any third-party package you can use the Go command-line tool to download the code.

### Go rules

- Go code is delivered in packages and you are free to use the functionality found in existing packages.

- You either use a variable or you do not declare it at all.

- There is only one way to format curly braces in Go.

- Coding blocks in Go are embedded in curly braces even if they contain just a single statement or no statements at all.

- Go functions can return multiple values.

- You cannot automatically convert between different data types, even if they are of the same kind.

## Go CLI

1. go build -> compiles

2. go run -> compiles and executes

3. go fmt -> format all code in each file in current directory

4. go install -> compiles and installs a package

5. go get -> download the raw source code of someone package

6. go test -> runs any tests associated with the current project

## Table of Content:

1. Pointers

2. Nil and Zero values

3. Init Functions

## Pointers

Go works with arguments as values or references. When working with
references we talk about pointers. A pointer addresses a memory location
instead of a value. In Go pointers are identified following the C notation with
a star. For a type T, *T indicates a pointer to a value of type T.

```go
package main

import "fmt"

func a (i int){
    i = 0
}

func b (i *int) {
    *i = 0
}

func main() {
    x := 100
    a(x)
    fmt.Println(x)

    b(&x)
    fmt.Println(x)

    fmt.Println(&x)
}
```

Notice that a does not change x value because it receives values as arguments. This is a works with a copy of variable x.

However, function b sets x to zero because it receives a pointer to the variable.

The operator & returns the pointer to the variable, which is of type *int. See how this operator returns the memory address of variable x with fmt.Println(&x).

How to decide when to use a pointer or a value depends on the use case. If a value is intended to be modified in different parts of the code, passing pointers seems reasonable.

## Nil and Zero Values

When a variable is created and not initialized, the compiler automatically assigns it a default value. This value depends on the variable type. The keyword nil specifies a particular value for every non-initialized type. Notice that nil is not an undefined value like in other programming languages, nil is a value itself.

```go
package main

import "fmt"


func main() {

    var a int
    fmt.Println(a)

    var b *int
    fmt.Println(b)

    var c bool
    fmt.Println(c)

    var d func()
    fmt.Println(d)

    var e string
    fmt.Printf("[%s]",e)
}
```

## Init Functions

We have mentioned that every program in Go must have a main package with a main function to be executed. However, this imposes some limitations for certain solutions such as libraries. Imagine we import a library into our code. A library is not designed to be executed, it offers data structures, methods, functions, etc. Libraries probably do not even have a main package. If this library requires some initial configuration before invoked (initialize variables, detect the operating system, etc.) it looks impossible.

Go defines init functions that are executed once per package. When we import a package the Go runtime follows this order:

1. Initialize imported packages recursively.

2. Initialize and assign values to variables.

3. Execute `init` functions.

```go
package main

import "fmt"

var x = xSetter()

func xSetter() int{
    fmt.Println("xSetter")
    return 42
}

func init() {
    fmt.Println("Init function")
}

func main() {
    fmt.Println("This is the main")
}
```

The init function has no arguments neither returns any value. A package can have several init functions and they cannot be invoked from any part of the code.

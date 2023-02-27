## Table of Content:

1. Errors, Panic, Recover
2. Packages, imports, and exports
3. Type casting
4. Type inference
5. Arrays
6. Slices
7. Maps
8. Allocation with make()
9. Structs

## Errors, Panic, Recover

All error handling operations in Go are based on the type error. An error variable stores a message with some information. In situations where an error can occur, the usual way to proceed is to return a filled error informing about its cause. This can be done using the errors.New function.

```go
type error interface {
    Error() string
}
```

Errors can be constructed on the fly using Go’s built-in `errors` or `fmt` packages. For example, the following function uses the `errors` package to return a new error with a static error message:

```go
package main

import "errors"

func DoSomething() error {
    return errors.New("something didn't work")
}
```

Similarly, the `fmt` package can be used to add dynamic data to the error, such as an `int`, `string`, or another `error`. For example:

```go
package main

import "fmt"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide '%d' by zero", a)
    }
    return a / b, nil
}
```

There are a few other important things to note in the example above:

- Errors can be returned as `nil`, and in fact, it’s the default, or “zero”, value of on error in Go. This is important since checking `if err != nil` is the idiomatic way to determine if an error was encountered (replacing the `try`/`catch` statements you may be familiar with in other programming languages).

- Errors are typically returned as the last argument in a function. Hence in our example above, we return an `int` and an `error`, in that order.

- When we do return an error, the other arguments returned by the function are typically returned as their default “zero” value. A user of a function may expect that if a non-nil error is returned, then the other arguments returned are not relevant.

- Lastly, error messages are usually written in lower-case and don’t end in punctuation. Exceptions can be made though, for example when including a proper noun, a function name that begins with a capital letter, etc.

eg with std package:

```go
func Open(name string) (file *File, err error)

f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

##### Defining Expected Errors

Defining expected Errors so they can be checked for explicitly in other parts of the code. This becomes useful when you need to execute a different branch of code if a certain kind of error is encountered.

###### Defining Sentinel Errors

```go
package main

import (
    "errors"
    "fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

func main() {
    a, b := 10, 0
    result, err := Divide(a, b)
    if err != nil {
        switch {
        case errors.Is(err, ErrDivideByZero):
            fmt.Println("divide by zero error")
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}
```

###### Defining Custom Error Types

```go
package main

import (
    "errors"
    "fmt"
)

type DivisionError struct {
    IntA int
    IntB int
    Msg  string
}

func (e *DivisionError) Error() string { 
    return e.Msg
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, &DivisionError{
            Msg: fmt.Sprintf("cannot divide '%d' by zero", a),
            IntA: a, IntB: b,
        }
    }
    return a / b, nil
}

func main() {
    a, b := 10, 0
    result, err := Divide(a, b)
    if err != nil {
        var divErr *DivisionError
        switch {
        case errors.As(err, &divErr):
            fmt.Printf("%d / %d is not mathematically valid: %s\n",
              divErr.IntA, divErr.IntB, divErr.Error())
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}
```

##### Defer, Panic, and Recover

**Defer**

The defer statement pushes a function onto a list. This list of functions is executed when the surrounding function ends. This statement is specially designed to ensure the correctness of the execution after the function ends. In particular, defer is useful to clean up resources allocated to a function.

```go
package main

import "fmt"

func CloseMsg() {
    fmt.Println("Closed!!!")
}

func main() {
    defer CloseMsg()
    fmt.Println("Doing something...")
    defer fmt.Println(("Certainly closed!!!"))
    fmt.Println("Doing something else...")
}
```

3 defer simple rule:

1. A defered function's arguments are evaluated when the defer statement is evaluated
   
   ```go
   func a() {
       i := 0
       defer fmt.Println(i)
       i++
       return
   }
   ```

2. Defered function calls are executed in Last In First Out order after the surrounding function returns
   
   ```go
   func b() {
       for i := 0; i < 4; i++ {
           defer fmt.Print(i)
       }      
   }
   ```

3. Deferred functions may read and assingn to the returning function's named return values
   
   ```go
   func c() (i int) {
       defer func() { i++ }()
       return 1
   }
   ```

**Panic**

Stops the execution flow, executes deferred functions and returns control to the calling function. This occurs for all functions until the program crashes. A call to panic indicates a situation that goes beyond the control of the program.

```go
package main

import "fmt"

func something() {
    defer fmt.Println("closed something")
    for i:=0;i<5;i++ {
        fmt.Println(i)
        if i> 2 {
            panic("Panic was called")
        }
    }
}

func main() {
    defer fmt.Println("closed main")
    something()
    fmt.Println("something was finished")
}
```

**Recover**

It may occur that under certain conditions when panic is invoked, the control flow can be restored. The recover built-in function used inside a deferred function can be used to resume normal execution.

Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

## Packages, imports, and exports

Go programs are organized into packages. A package is a group of one or more source files which code is accessible from the same package. Additionally, a package can be exported and used in other packages.

The package main is a special case that informs the Go compiler to consider that package as the entry point for an executable file. Actually, the package main is expected to have a main function in order to be compiled.

###### Import paths

The import path of packages is globally unique. To avoid conflict between the path of the packages other than the standard library, the package path should start with the internet domain name of the organization that owns or host the package.

```go
import "fmt"
import "geeksforgeeks.com/example/strings"
```

###### Package declaration

package declaration is always present at the beginning of the source file and the purpose of this declaration is to determine the default identifier for that package when it is imported by another package.

```go
package main
```

###### Import declaration

The import declaration immediately comes after the package declaration. The Go source file contains zero or more import declaration and each import declaration specifies the path of one or more packages in the parentheses.

```go
// Importing single package
import "fmt"

// Importing multiple packages
import(
"fmt"
"strings"
"bytes"
) 
```

###### Blank import

sometimes we import some packages in our program, but we do not use them in our program.

```go
import _ "strings"
```

###### Nested packages

 you are allowed to create a package inside another package simply by creating a subdirectory. And the nested package can import just like the root package.

```go
import "math/cmplx"
```

###### Giving names to the packages

when you name a package you must always follow the following points:

- When you create a package the name of the package must be short and simple. For example strings, time, flag, etc. are standard library package.
- The package name should be descriptive and unambiguous.
- Always try to avoid choosing names that are commonly used or used for local relative variables.
- The name of the package generally in the singular form. Sometimes some packages named in plural form like strings, bytes, buffers, etc. Because to avoid conflicts with the keywords.
- Always avoid package names that already have other connotations.

managing dependency: [Managing dependencies - The Go Programming Language](https://go.dev/doc/modules/managing-dependencies)

## Type casting

Type-casting means converting one type to another. Any type can be converted to another type but that does not guarantees that the value will remain intact or in fact preserved at all as we will see in this post.

Statically typed languages like [C](https://www.geeksforgeeks.org/c-programming-language/)/[C++](https://www.geeksforgeeks.org/c-plus-plus/), [Java](https://www.geeksforgeeks.org/java/), provide the support for Implicit Type Conversion but Golang is different, as it **doesn’t support the** ***Automatic Type Conversion or Implicit Type Conversion*** even if the data types are compatible. The reason for this is the Strong Type System of the Golang which doesn’t allow to do this. For type conversion, you must perform explicit conversion.

As per [Golang Specification](https://golang.org/ref/spec), there is no typecasting word or terminology in Golang. If you will try to search Type Casting in Golang Specifications or Documentation, you will find nothing like this. There is only Type Conversion.

```go
package main

import (
    "fmt"
)

func main() {
    f := 12.34567
    i := int(f)  // loses precision
    fmt.Println(i)      // 12

    ii := 34
    ff := float64(ii)

    fmt.Println(ff)     // 34
}
```

**Strings and bytes conversion**

```go
package main

import (
    "fmt"
)

func main() {
    var s string = "Hello World"
    var b []byte = []byte(s)     // convert ty bytes

    fmt.Println(b)  // [72 101 108 108 111 32 87 111 114 108 100]

    ss := string(b)              // convert to string

    fmt.Println(ss)     // Hello World
}
```

## Type inference

Although Go is a Statically typed language, It doesn’t require you to explicitly specify the type of every variable you declare. When you declare a variable with an initial value, Golang automatically infers the type of the variable from the value on the right-hand side.

```go
package main

import "fmt"

func main() {
    v := 42.5 // change me!
    fmt.Printf("v is of type %T\n", v)
}
```

## Arrays

By definition, an array is an indexed sequence of elements with a given length. Like any other Go variable, arrays are typed and their size is fixed. By default, this array is filled with zeros. Every array has a len function that returns the array length.

```go
package main

import (
    "fmt"
)

func main() {
    var a[5] int
    fmt.Println(a)

    b := [5]int{0,1,2,3,4}
    fmt.Println(b)

    c := [5]int{0,1,2}
    fmt.Println(c)

}
```

There are major differences between the ways arrays work in Go and C. In Go:

- Arrays are values. Assigning one array to another copies all the elements.
- In particular, if you pass an array to a function, it will receive a *copy* of the array, not a pointer to it.
- The size of an array is part of its type. The types `[10]int` and `[20]int` are distinct.

## Slices

Descriptor for a contiguous segment of an underlying array and provides access to a
numbered sequence elements from that array. In other words, a slice is a reference to an array. The slice itself does not store any data but offers a view of it.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]string{"a","b","c","d","e"}
    fmt.Println(a)
    fmt.Println(a[:])
    fmt.Println(a[0])
    fmt.Println(a[0],a[1],a[2],a[3],a[4])
    fmt.Println(a[0:2])
    fmt.Println(a[1:4])
    fmt.Println(a[:2])
    fmt.Println(a[2:])
}
```

different output:

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := [5]string{"a","b","c","d","e"}

    fmt.Println(reflect.TypeOf(a))
    fmt.Println(reflect.TypeOf(a[0:3]))
    fmt.Println(reflect.TypeOf(a[0]))
}
```

###### Length and capacity

An important difference between an array and a slice is the concept of capacity. While an array allocates a fixed amount of memory that is directly related to its length, a slice can reserve a larger amount of memory that does not necessarily have to be filled. The filled memory corresponds to its length, and all the available memory is the capacity. Both values are accessible using functions len and cap.

```go
package main

import (
    "fmt"
)

func main() {
    a := []int{0,1,2,3,4}
    fmt.Println(a, len(a), cap(a))

    b := append(a,5)
    fmt.Println(b, len(b), cap(b))

    b = append(b,6)
    fmt.Println(b, len(b), cap(b))

    c := b[1:4]
    fmt.Println(c, len(c), cap(c))

    d := make([]int,5,10)
    fmt.Println(d, len(d), cap(d))
    // d[6]=5 —> This will fail
}
```

###### Iteration

The most common operation you can find in a slice is the iteration through its items. Any for loop is a good candidate to do this. Go simplifies iterations through collections with the range clause to.

```go
package main

import (
    "fmt"
)

func main() {
    names := []string{"Jeremy", "John", "Joseph",}

    for i:=0;i<len(names);i++{
        fmt.Println(i,names[i])
    }

    for position, name := range names {
        fmt.Println(position,name)
    }

}
```

A correct approach to modify the iterated slice is to access the original variable with the corresponding index.

```go
package main

import (
    "fmt"
)

func main() {

    names := []string{"Jeremy", "John", "Joseph"}
    for _, name := range(names){
        // name is a copy
        name = name + "_changed"
    }
    fmt.Println(names)

    for position, name := range(names){
        // this modifies the original value
        names[position] = name + "_changed"
    }
    fmt.Println(names)
}
```

## Maps

A map is a construct that maps a key with a value. Keys are intended to be unique and can be of any type that implements == and != operators.

```go
package main

import (
    "fmt"
)

func main() {

    var ages map[string]int
    fmt.Println(ages)
    // This fails, ages was not initialized
    // ages["Jesus"] = 55

    ages = make(map[string]int,5)
    // Now it works because it was initialized
    ages["Jesus"] = 33
    ages = map[string]int{
        "Jesus": 33,
        "Mathusalem": 969,
    }
    fmt.Println(ages)

}
```

```go
package main

import (
    "fmt"
)

func main() {

    birthdays := map[string]string{
        "Jesus": "12-25-0000",
        "Budha": "563 BEC",
    }
    fmt.Println(birthdays,len(birthdays))

    xmas, found := birthdays["Jesus"]
    fmt.Println(xmas, found)
    delete(birthdays, "Jesus")

    fmt.Println(birthdays,len(birthdays))
    _, found = birthdays["Jesus"]
    fmt.Println("Did we find when its Xmas?", found)

    birthdays["Jesus"]="12-25-0000"
    fmt.Println(birthdays)
}
```

###### Iterate

To iterate a map we would require the collection of keys. Fortunately, the range built-in function offers a simple solution to iterate through all the key-value pair of any map. The rules explained for slices apply in the case of maps. For every iteration range returns the current key and value.

```go
package main

import (
    "fmt"
)

func main() {

    sales := map[string]int {
        "Jan": 34345,
        "Feb": 11823,
        "Mar": 8838,
        "Apr": 33,
    }

    fmt.Println("Month\tSales")
    for month, sale := range sales {
        fmt.Printf("%s\t\t%d\n",month,sale)
    }

}
```

## Allocation with make

## Struct

A sequence of elements named fields. Each field has a name and a type.

```go
package main

import (
	"fmt"
)

type Rectangle struct{
    Height int 
    Width int
}

func main() {
    a := Rectangle{}
    fmt.Println(a)

    b := Rectangle{4,4}
    fmt.Println(b)

    c := Rectangle{Width: 10, Height: 3}
    fmt.Println(c)

    d := Rectangle{Width: 7}
    fmt.Println(d)
}
```

Using function to ensure required parameter and added error handling

```go
package main

import (
	"errors"
	"fmt"
)

type Rectangle struct{
    Height int 
    Width int
}

func NewRectangle(height int, width int) (*Rectangle, error) {
    if height <= 0 || width <= 0 {
        return nil, errors.New("params must be greater than zero")
    }
    return &Rectangle{height, width}, nil
}
    

func main() {
    a := Rectangle{Height: 7}
    fmt.Println(a)

    r, err := NewRectangle(1,3)
    
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(r)
    // fmt.Println(*r)
}
```

###### Anonymous structs

Compared with a regular struct like Circle printing the struct brings a similar result. However, we cannot print its name as we do with type Circle. The fields from the anonymous function can be modified like done with regular structs. Notice that these anonymous structures can be compared with other structures if and only if they have the same fields.

```go
package main

import (
	"fmt"
	"reflect"
)

type Circle struct{
    x int 
    y int
    radius int
}

func main() {
    ac := struct{x int; y int; radius int}{1,2,3}
     c := Circle{10,10,3}
    
     fmt.Printf("%+v\n", ac)
     fmt.Println(reflect.TypeOf(ac))
     fmt.Printf("%+v\n",c)
     fmt.Println(reflect.TypeOf(c))
    
     ac.x=3
     fmt.Printf("%+v\n", ac)
    
     ac = c
     fmt.Printf("%+v\n", ac)
     fmt.Println(reflect.TypeOf(ac))
}
```

###### Nested structs

Structs can be nested to incorporate other structs definition.

```go
package main

import (
	"fmt"
)

type Coordinates struct{
    x int
    y int
}

type Circle struct{
    center Coordinates
    radius int
}

func main() {
    c := Circle{Coordinates{1, 2}, 3}
    fmt.Printf("%+v\n", c)
}
```

###### Embedded structs

To embed a struct in other structs, this has to be declared as a nameless field.

```go
package main

import (
	"fmt"
)

type A struct { fieldA int }

type B struct { fieldA int }

type C struct {
    A
    B
}
    

func main() {
    // Embedding structs can be done only if the compiler find no ambiguities

    a := A{10}
    b := B{20}
    c := C{a,b}

    // —> Ambiguos access
    // fmt.Println(c.fieldA)
    fmt.Println(c.A.fieldA,c.B.fieldA)
}
```

# Table of Content

- Numeric data types

- Non-numeric data types

- Go constants

- Composite types
  
  - Non-Reference Types
    
    - Arrays
    
    - Structs
  
  - Reference Types
    
    - Slices
    
    - Channels
    
    - Maps
    
    - Pointers
    
    - Functions
  
  - Interface
    
    - Special case of empty interface

## Numeric data types

| **Data Type** | **Description**                   |
| ------------- | --------------------------------- |
| int8          | 8-bit signed integer              |
| int16         | 16-bit signed integer             |
| int32         | 32-bit signed integer             |
| int64         | 64-bit signed integer             |
| int           | 32- or 64-bit signed integer      |
| uint8         | 8-bit unsigned integer            |
| uint16        | 16-bit unsigned integer           |
| uint32        | 32-bit unsigned integer           |
| uint64        | 64-bit unsigned integer           |
| uint          | 32- or 64-bit unsigned integer    |
| float32       | 32-bit floating-poing number      |
| float64       | 64-bit floating-poing number      |
| complex64     | Complex number with float32 parts |
| complex128    | Complex number with float64 parts |

```go
package main

import (
    "fmt"
    "math/bits"
    "reflect"
    "unsafe"
)

func main() {
    //This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
    sizeOfIntInBits := bits.UintSize
    fmt.Printf("%d bits\n", sizeOfIntInBits)

    var a int
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

    b := 2
    fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))
}
```

- **bits** package of golang can help know the size of an **int** on your system
- **unsafe.Sizeof()** function can also be used to see the size of int in bytes

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    //Declare a int 8
    var a int8 = 2

    //Size of int8 in bytes
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
```

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    //Declare a uint8

    var a uint8 = 2

    //Size of uint8 in bytes
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
```

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    //Declare a float32
    var a float32 = 2

    //Size of float32 in bytes
    fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
    fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
```

```go
package main
import (
    "fmt"
    "reflect"
    "unsafe"
)
func main() {
    var a float32 = 3
    var b float32 = 5

    //Initialize-1
    c := complex(a, b)

    //Initialize-2
    var d complex64
    d = 4 + 5i

    //Print Size
    fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c))
    fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d))

    //Print type
    fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
    fmt.Printf("d's type is %s\n", reflect.TypeOf(d))

    //Operations on complex number
    fmt.Println(c+d, c-d, c*d, c/d)
}
```

## Non-numeric data types

###### Bytes

byte in Go is an alias for **uint8** meaning it is an integer value. This integer value is of 8 bits and it represents one byte i.e number between 0-255). A single byte therefore can represent ASCII characters. Golang does not have any data type of ‘char’. Therefore

- byte is used to represent the ASCII character

- rune is used to represent all UNICODE characters which include every character that exists. We will study about rune later in this tutorial.

```go
package main
import (
    "fmt"
    "reflect"
    "unsafe"
)
func main() {
    var r byte = 'a'

    //Print Size
    fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

    //Print Type
    fmt.Printf("Type: %s\n", reflect.TypeOf(r))

    //Print Character
    fmt.Printf("Character: %c\n", r)
    s := "abc"

    //This will the decimal value of byte
    fmt.Println([]byte(s))
}
```

###### Runes

rune in Go is  an alias for **int32** meaning it is an integer value. This integer value is meant to represent a Unicode Code Point. To understand rune you have to know what Unicode is. Below is short description but you can refer to famous blog post about it – [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](http://www.joelonsoftware.com/articles/Unicode.html)

Unicode => Unicode is a superset of ASCII characters which assigns a unique number to every character that exists. This unique number is called Unicode Code Point.

eg:

- Digit **0** is represented as Unicode Point **U+0030 (Decimal Value – 48)**

- Small Case **b** is represented as Unicode Point  **U+0062 (Decimal Value – 98)**

- A pound symbol **£** is represented as Unicode Point **U+00A3 (Decimal Value – 163)**

###### utf-8

utf-8 saves every Unicode Point either using 1, 2, 3 or 4 bytes. ASCII points are stored using 1 byte. That is why rune is an alias for int32 because a Unicode Point can be of max 4 bytes in Go as in GO every string is encoded using utf-8.

```go
package main
import (
    "fmt"
    "reflect"
    "unsafe"
)
func main() {
    r := 'a'

    //Print Size
    fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

    //Print Type
    fmt.Printf("Type: %s\n", reflect.TypeOf(r))

    //Print Code Point
    fmt.Printf("Unicode CodePoint: %U\n", r)

    //Print Character
    fmt.Printf("Character: %c\n", r)
    s := "0b£"

    //This will print the Unicode Points
    fmt.Printf("%U\n", []rune(s))

    //This will the decimal value of Unicode Code Point
    fmt.Println([]rune(s))
}
```

###### String

string is a read only slice of bytes in golang. String can be initialized in two ways. eg:

- using double quotes -> "this”

- using back quotes -> \\`this`

```go
package main
import (
    "fmt"
)
func main() {
    //String in double quotes
    x := "this\nthat"
    fmt.Printf("x is: %s\n", x)

    //String in back quotes
    y := `this\nthat`
    fmt.Printf("y is: %s\n", y)
    s := "ab£"

    //This will print the byte sequence. 
    //Since character a and b occupies 1 byte each and £ character occupies 2 bytes. 
    //The final output will 4 bytes
    fmt.Println([]byte(s))

    //The output will be 4 for same reason as above
    fmt.Println(len(s))

    //range loops over sequences of byte which form each character
    for _, c := range s {
        fmt.Println(string(c))
    }

    //Concatenation
    fmt.Println("c" + "d")
}
```

###### Booleans

The data type is **bool** and has two possible values true or false.

Operations:

- AND – &&
- OR  – ||
- Negation – !

```go
package main

import "fmt"

func main() {
    //Default value will be false it not initialized
    var a bool
    fmt.Printf("a's value is %t\n", a)

    //And operation on one true and other false
    andOperation := 1 < 2 && 1 > 3
    fmt.Printf("Ouput of AND operation on one true and other false %t\n", andOperation)

    //OR operation on one true and other false
    orOperation := 1 < 2 || 1 > 3
    fmt.Printf("Ouput of OR operation on one true and other false: %t\n", orOperation)

    //Negation Operation on a false value
    negationOperation := !(1 > 2)
    fmt.Printf("Ouput of NEGATION operation on false value: %t\n", negationOperation)
}
```

###### Times and dates

The king of working with times and dates in Go is the time.Time data type, which  represents an instant in time with nanosecond precision. Each time.Time value is  associated with a location (time zone). If you are a UNIX person, you might already know about the UNIX epoch time and wonder how to get it in Go. The time.Now().Unix() function returns the popular UNIX epoch time, which is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970. If you want to convert the UNIX time to the equivalent time. Time value, you can use the time.Unix() function.

The documentation of the time package (https://golang.org/pkg/time/) contains 
even more detailed information about parsing dates and times.

```go
package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    start := time.Now()
    if len(os.Args) != 2 {
        fmt.Println("Usage: dates parse_string")
        return
    }
    dateString := os.Args[1]

    // Is this a date only?
    d, err := time.Parse("02 January 2006", dateString)
    if err == nil {
        fmt.Println("Full:", d)
        fmt.Println("Time:", d.Day(), d.Month(), d.Year())
    }

    // Is this a date + time?
    d, err = time.Parse("02 January 2006 15:04", dateString)
    if err == nil {
        fmt.Println("Full:", d)
        fmt.Println("Date:", d.Day(), d.Month(), d.Year())
        fmt.Println("Time:", d.Hour(), d.Minute())
    }

    // Is this a date + time with month represented as a number?
    d, err = time.Parse("02-01-2006 15:04", dateString)
    if err == nil {
        fmt.Println("Full:", d)
        fmt.Println("Date:", d.Day(), d.Month(), d.Year())
        fmt.Println("Time:", d.Hour(), d.Minute())
    }

   // Is it time only?
    d, err = time.Parse("15:04", dateString)
    if err == nil {
        fmt.Println("Full:", d)
        fmt.Println("Time:", d.Hour(), d.Minute())
    }
    t := time.Now().Unix()
    fmt.Println("Epoch time:", t)

    // Convert Epoch time to time.Time
    d = time.Unix(t, 0)
    fmt.Println("Date:", d.Day(), d.Month(), d.Year())
    fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
    duration := time.Since(start)
    fmt.Println("Execution time:", duration)
}
```

###### Working with different timezone

This can be particularly handy when you want to preprocess log files from different sources that use different time zones in order to convert these different time zones into a common one. 

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    loc, _ := time.LoadLocation("America/New_York")
    fmt.Printf("New York Time: %s\n", now.In(loc))
}
```

##### Non-Reference Types

###### Arrays

Arrays in go are values. They are fixed-length sequences of the same type

- When you assign an array to another variable, it copies the entire array

- When you pass an array as an argument to a function, it makes an entire copy of the array instead of passing just the address

```go
newArray := [n]Type{val1, val2, val3}
newArray := [len]Type{}
```

```go
package main

import "fmt"

func main() {
    //Declare a array
    sample := [3]string{"a", "b", "c"}
    print(sample)
}

func print(sample [3]string) {
    fmt.Println(sample)
}
```

###### Struct

In GO struct is named collection of fields. These fields can be of different types. Struct acts as a container of related data of heterogeneous data type.

- Name of string type

- Age of int type

- DOB of time.Time type

```go
type employee struct {
    name string
    age  int
    dob  time.Time
}
```

```go
package main
import (
    "fmt"
)
//Declare a struct
type employee struct {
    name   string
    age    int
    salary float64
}
func main() {
    //Initialize a struct without named fields
    employee1 := employee{"John", 21, 1000}
    fmt.Println(employee1)

    //Initialize a struct with named fields
    employee2 := employee{
        name:   "Sam",
        age:    22,
        salary: 1100,
    }
    fmt.Println(employee2)

    //Initializing only some fields. Other values are initialized to default zero value of that type
    employee3 := employee{name: "Tina", age: 24}
    fmt.Println(employee3)
}
```

##### Reference Types

###### Slices

Slices are dynamically sized, reference into the elements of an array.

- Address to the underlying array

- Length of the slice

- Capacity of the slice

```go
make([]TYPE, length, capacity)
p := []string{"a", "b", "c"}
```

```go
package main

import "fmt"

func main() {
    //Declare a slice using make
    s := make([]string, 2, 3)
    fmt.Println(s)

    //Direct intialization
    p := []string{"a", "b", "c"}
    fmt.Println(p)

    //Append function
    p = append(p, "d")
    fmt.Println(p)

    //Iterate over a slcie
    for _, val := range p {
        fmt.Println(val)
    }
}
```

###### Channels

Channels provide synchronization and communication between goroutines. You can think of it as a pipe through which goroutines can send values and receive values. The operation <- is used to send or receive, with direction of arrow specifying the direction of flow of data

```go
ch <- val    //Sending a value present in var variable to channel
val := <-cha  //Receive a value from  the channel and assign it to val variable
```

Channel are of two types

- **Unbuffered Channel**- It doesn't have any capacity to hold and values and thus
  
  - Send on a channel is block unless there is another goroutine to receive.
  - Receive is block until there is another goroutine on the other side to send.

- **Buffered Channel-** You can specify the size of buffer here and for them
  
  - Send on a buffer channel only blocks if the buffer is full
  - Receive is the only block is buffer of the channel is empty

**Bufferd Channel**

```go
package main

import "fmt"

func main() {
    //Creating a buffered channel of length 3
    eventsChan := make(chan string, 3)
    eventsChan <- "a"
    eventsChan <- "b"
    eventsChan <- "c"
    //Closing the channel
    close(eventsChan)
    for event := range eventsChan {
        fmt.Println(event)
    }
}
```

**UnBufferd Channel**

```go
package main

import "fmt"

func main() {
    eventsChan := make(chan string)
    go sendEvents(eventsChan)
    for event := range eventsChan {
        fmt.Println(event)
    }
}

func sendEvents(eventsChan chan<- string) {
    eventsChan <- "a"
    eventsChan <- "b"
    eventsChan <- "c"
    close(eventsChan)
}
```

###### Maps

maps are golang builtin datatype similar to a hash which map key to a value. maps are referenced data types. When you assign one map to another both refer to the same underlying map.

```go
var employeeSalary map[string]int
var employeeSalary make(map[string]int)

//Empty braces
employeeSalary := map[string]int{}

//Specify values
employeeSalary := map[string]int{
"John": 1000
"Sam": 2000
} 

// operations
employeeSalary["John"] = 1000
salary := employeeSalary["John"]
delete(employeeSalary, "John")
```

```go
package main

import "fmt"

func main() {
    //Declare
    var employeeSalary map[string]int
    fmt.Println(employeeSalary)

    //Intialize using make
    employeeSalary2 := make(map[string]int)
    fmt.Println(employeeSalary2)

    //Intialize using map lieteral
    employeeSalary3 := map[string]int{
        "John": 1000,
        "Sam":  1200,
    }
    fmt.Println(employeeSalary3)

    //Operations
    //Add
    employeeSalary3["Carl"] = 1500

    //Get
    fmt.Printf("John salary is %d\n", employeeSalary3["John"])

    //Delete
    delete(employeeSalary3, "Carl")

    //Print map
    fmt.Println("\nPrinting employeeSalary3 map")
    fmt.Println(employeeSalary3)
}
```

###### Pointers

Pointer is a variable that holds a memory address of another variable. The zero value of a pointer is nil.

```go
var ex *int
// initialize
a := 2
b := &b // $ used to get the address of a variable
fmt.Println(*b) // deference a pointer which means gettine the value at address stored in pointer

// can also be initialize using new operator
a := new(int)
*a = 100
fmt.Printltn(*a) // output will be 10
```

```go
package main

import "fmt"

func main() {
    //Declare
    var b *int
    a := 2
    b = &a
    
    //Will print a address. Output will be different everytime.
    fmt.Println(b)
    fmt.Println(*b)
    b = new(int)
    *b = 10
    fmt.Println(*b) 
}
```

###### Function/Methods

In Go function are values and can be passed around like a value. Basically, function can be used as first-order objects and can be passed around.  The signature of a function is

```go
func some_func_name(arguments) return_values
```

```go
package main

import "fmt"

func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1, 2))
}

func doOperation(fn func(int, int) int, x, y int) int {
    return fn(x, y)
}
```

##### Interface

Interface is a type in Go which is a collection of method signatures. Any type which implements all methods of the interface is of that interface type. Zero value of an interface is nil.

```go
type name_of_interface interface{
//Method signature 1
//Method signature 2
}
```

There is no explicit declaration that a type implements an interface. In fact, in Go there doesn't exist any **"implements"** keyword similar to Java.  A type implements an interface if it implements all the methods of the interface.

```go
package main

import "fmt"

type shape interface {
    area() int
}

type square struct {
    side int
}

func (s *square) area() int {
    return s.side * s.side
}

func main() {
    var s shape
    s = &square{side: 4}
    fmt.Println(s.area())
}
```

###### Special case of empty interface

An empty interface has no methods, hence by default all concrete types implement the empty interface. If you write a function that accepts an empty interface then you can pass any type to that function.

```go
package main

import "fmt"

func main() {
    test("thisisstring")
    test("10")
    test(true)
}

func test(a interface{}) {
    fmt.Printf("(%v, %T)\n", a, a)
}
```

ref: [All data types in Golang with examples - Welcome To Golang By Example](https://golangbyexample.com/all-data-types-in-golang-with-examples/#Integers_Signed_and_UnSigned)



## Table of Content:

1. Basic syntax

2. variables and declaration

3. Conditionals(if, switch statements)

4. Iterating(loops and range)

5. Functions, multiple/named returns

## Defining & using variables

Go provides multiple ways to declare new variables in order to make the variable declaration process more natural and convenient.

```go
package main

import (
    "fmt"
    "math"
)

var Global int = 1234
var AnotherGlobal = -5678
var c, python, java bool
// with initializer
// var i, j int = 1, 2
// var c, python, java = true, false, "no!"

// constant
const (
    Pi = 3.14
)

func main() {
    var i int
    // short variable declaration
    k:= 3
    fmt.Println(i, c, python, java)
    l:= math.Abs(float64(AnotherGlobal))
    fmt.Printf("Global=%d, l=%.2f.\n", Global, l)
}
```

- A second global variable named AnotherGlobal—Go automatically infers its data type from its value, which in this case is an integer

- As math.Abs() requires a float64 parameter, you cannot pass AnotherGlobal to it because AnotherGlobal is an int variable. The float64() type cast converts the value of AnotherGlobal to float64. Note that AnotherGlobal continues to be int.

- fmt.Printf() formats and prints our output.

## 

## Conditionals(if, switch statements)

###### if Else

```go
// Basic syntax
if x > max {
    x = max
}


if x <= y {
    min = x
} else {
    min = y
}

// With init statement
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else {
    fmt.Println(num, "is positive")
}

// Ternary operator
res = expr ? x : y
```

example of a common situation where code must guard against a sequence of error conditions.

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

###### Switch statements

There is no automatic fall through in go, but cases can be presented in comma-separated lists.

```ts
package main

import (
    "fmt"
    "time"
)

func main() {
    today := time.Now()

    switch today.Day() {
    case 5:
        fmt.Println("Today is 5th. Clean your house.")
    case 10:
        fmt.Println("Today is 10th. Buy some wine.")
    case 15:
        fmt.Println("Today is 15th. Visit a doctor.")
    // fallthrough -> force to fall through the succesive case block
    case 25:
        fmt.Println("Today is 25th. Buy some food.")
        fallthrough
    // multiple casses statements
    case 28, 29, 30:
        fmt.Println("Party tonight.")
    default:
        fmt.Println("No information available for that day.")
    }
}
```

Sometimes, though, it's necessary to break out of a surrounding loop, not the switch, and in Go that can be accomplished by putting a label on the loop and "breaking" to that label. This example shows both uses.

```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

###### Type switch

A switch can also be used to discover the dynamic type of an interface variable. Such a *type switch* uses the syntax of a type assertion with the keyword `type` inside the parentheses.

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## 

## Iterating (loops and range)

Instead of including direct support for while loops. However, depending on how you write a for loop, it can function as a while loop or an infinite loop. Moreover, for loops can implement the functionality of JavaScript's forEach function when combined with the range keyword.

A for loop can be exited
with a break keyword and you can skip the current iteration with the continue keyword. When used with range, for loops allow you to visit all the elements of a slice or an array without knowing the size of the data structure.

```go
package main
import "fmt"
func main() {
    // Traditional for loop
    for i := 0; i < 10; i++ {
        fmt.Print(i*i, " ")
    }
    fmt.Println()
}
```

idiomatic go:

```go
    i := 0
    for ok := true; ok; ok = (i != 10) {
        fmt.Print(i*i, " ")
        i++
    }
    fmt.Println()
```

how a for loop can simulate a while loop:

```go
// For loop used as while loop
i := 0
for {
    if i == 10 {
        break
    }
    fmt.Print(i*i, " ")
    i++
}
fmt.Println()

// or
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)


// A break statement leaves the innermost for, switch or select statement.
// A continue statement begins the next iteration of the innermost for loop at its post statement (i++).
```

range loop:

```go
// Looping over elements in slices, arrays, maps, channels or strings
aSlice := []int{-1, 2, 1, -1, 2, -2}
for i, v := range aSlice     {
    fmt.Println("index:", i, "value: ", v)
}

// if only second value needed:
sum := 0
for _, value := range array {
    sum += value
}
```

## Functions, multiple/named returns

A function encapsulates a piece of code that performs certain operations or logic that is going to be required by other sections of the code. A function is the most basic solution to reuse code.

A function receives none or several parameters and returns none or several values. Functions are defined by keyword func, the arguments with their types, and the types of the returned values.

Two arguments function

```go
package main

import "fmt"

func sum(a int, b int) int {
    return a + b
}

func main() {
    result := sum(1, 2)
    fmt.Println(result)
}
```

Function returning several values

```go
package main

import "fmt"

func ops(a int, b int) (int,int) {
    return a + b, a - b
}

func main() {
    sum, subs := ops(2, 2)
    fmt.Println("2=2=",sum , "2-2=",subs)
    b, _ := ops(10,2)
    fmt.Println("10+2=",b)
}
```

Variadic function

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0

    for _, a := range(nums) {
        total = total + a
    }
    return total
}

func main() {
    total := sum(1,2,3,4,5)
    fmt.Println("The first five numbers sum is",total)
}
```

Functions as arguments

```go
package main

import "fmt"

func doit(operator func(int,int) int, a int, b int) int {
    return operator(a,b)
}

func sum(a int, b int) int {
    return a + b
}

func multiply(a int, b int) int {
    return a * b
}

func main() {
    c := doit(sum, 2, 3)
    fmt.Println("2+3=", c)
    d := doit(multiply, 2, 3)
    fmt.Println("2*3=", d)
}

//Higher order Function: functions that operate on other functions, either by taking them as arguments or by returning them.

package main

import "fmt"

func sum(x, y int) int {
    return x + y
}
func partialSum(x int) func(int) int {
    return func(y int) int {
        return sum(x, y)
    }
}
func main() {
    partial := partialSum(3)
    fmt.Println(partial(7))
}
```

Function closure

```go
package main

import "fmt"

func accumulator(increment int) func() int {
    i:=0
    return func() int {
        i=i+increment
        return i
    }
}

func main() {
    a:= accumulator(1)
    b:= accumulator(2)

    fmt.Println("a","b")
    for i:=0;i<5;i++ {
        fmt.Println(a(),b())
    }
}
```

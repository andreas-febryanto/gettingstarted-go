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
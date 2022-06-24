package main

import "fmt"

type Pair[A, B any] struct{}

func main() {
    const str = `a
b
c
`
    fmt.Println(str)
}

package main

import (
    "fmt"
    g "luvx/api/greet"
    "luvx/api/net"
)

func main() {
    s, i := g.Add(1, 3)
    fmt.Println(s, i)
    s, i = g.Times(2, 3)
    fmt.Println(s, i)

    //s, i = Add1(1, 3)
    //fmt.Println(s, i)
    //a()

    net.StartHttpServer()
}

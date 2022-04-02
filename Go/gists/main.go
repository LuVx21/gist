package main

import (
    "fmt"
    _ "luvx/api/log"
    "path"
)

func main() {
    logfile := path.Join("logDir", "app")
    fmt.Println(logfile)

    panic("hello")
}

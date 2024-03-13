package main

import (
    "fmt"
    "log"
    "net/http"
)

func pool(w http.ResponseWriter, r *http.Request) {
    fmt.Println("finish")
    fmt.Fprintln(w, "finish")
}

func main() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

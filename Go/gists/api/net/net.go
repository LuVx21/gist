package net

import (
    "fmt"
    "log"
    "net/http"
)

func StartHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func pool(w http.ResponseWriter, r *http.Request) {
    fmt.Println("finish")
    fmt.Fprintln(w, "finish")
}

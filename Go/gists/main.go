package main

import (
    "luvx/gin"
)

type Pair[A, B any] struct{}

func main() {
    gin.WebStart()
}

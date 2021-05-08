package greet

import (
    "fmt"
    "luvx/api/common"
)

type User common.User

func Add(a int, b int) (string, int) {
    return fmt.Sprintf("%d + %d =", a, b), a + b
}

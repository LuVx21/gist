package greet

import (
    "fmt"
)

func Times(a int, b int) (string, int) {
    return fmt.Sprintf("%d * %d =", a, b), a * b
}

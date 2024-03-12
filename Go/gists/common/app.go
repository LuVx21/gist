package common

import (
    "golang.org/x/exp/constraints"
    "log"
)

func IfThen[T constraints.Ordered](expr bool, a T, b T) T {
    if expr {
        return a
    }
    return b
}

// Catching 捕捉异常,避免异常退出
func Catching(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Print(r)
            }
        }()
        fn()
    }()
}

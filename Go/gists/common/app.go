package common

import "golang.org/x/exp/constraints"

func IfThen[T constraints.Ordered](expr bool, a T, b T) T {
    if expr {
        return a
    }
    return b
}

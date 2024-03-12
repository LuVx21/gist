package common

import "testing"

func Test_m1(t *testing.T) {
    Catching(task)
}

func task() {
    panic("异常")
}

package greet

import (
    "fmt"
    "testing"
)

func TestAdd(t *testing.T) {
    fmt.Println(Add(1, 3))
}

func TestAdd1(t *testing.T) {
    var a = User{1, "foo", 18}
    //fmt.Println(a, a.String())
    if &a != nil {
        t.Name()
        //t.Error("add(1, 2) should be equal to 3")
    }
    array := []int{1, 2, 3, 4}
    fmt.Println(array)
}

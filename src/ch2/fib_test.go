package fib

import (
    "testing"
)

func TestFibList(t *testing.T) {
    var (
        a int = 1
        b int = 1
        s int
    )

    for i := 0; i < 5; i++ {
        s, a, b = a, b, a+b
        t.Log(s)
    }
}

func TestCompare(t *testing.T) {
    a := [...]int{1, 2, 3, 4}
    b := [...]int{1, 2, 3, 4}
    c := [...]int{1, 2, 3, 5}
    t.Log(a == b)
    t.Log(a == c)
}

/**
按位清0
右边为1,直接清0
右边为0,左边是什么就是什么
*/
func TestRestZero(t *testing.T) {
    t.Log(0 &^ 0)
    t.Log(0 &^ 1)
    t.Log(1 &^ 1)
    t.Log(1 &^ 0)
}

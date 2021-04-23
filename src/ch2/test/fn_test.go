package try_test

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

func returnMultiValues() (int, int) {
    return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
    return func(n int) int {
        start := time.Now()
        ret := inner(n)

        fmt.Println("time spent: ", time.Since(start).Seconds())
        return ret
    }
}

func slowFun(op int) int {
    time.Sleep(time.Second * 1)
    return op + 1
}

func TestFn(t *testing.T) {
    a, b := returnMultiValues()
    t.Log(a, b)

    ts := timeSpent(slowFun)
    r := ts(10)
    t.Log(r)
}

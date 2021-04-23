package singleton

import (
    "fmt"
    "runtime"
    "sync"
    "testing"
    "time"
)

func runTask(id int) string {
    time.Sleep(time.Millisecond * 10)
    return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
    numOfRunner := 10
    ch := make(chan string, numOfRunner)
    for i := 0; i < numOfRunner; i++ {
        go func(i int) {
            ret := runTask(i)
            ch <- ret
        }(i)
    }
    return <-ch
}

func AllResponse() string {
    numOfRunner := 10
    ch := make(chan string, numOfRunner)
    for i := 0; i < numOfRunner; i++ {
        go func(i int) {
            ret := runTask(i)
            ch <- ret
        }(i)
    }

    finalRet := ""
    for j := 0; j < numOfRunner; j++ {
        finalRet += (<-ch) + "\n"
    }
    return finalRet
}

func TestFirstResponse(t *testing.T) {
    t.Log("Before: ", runtime.NumGoroutine())
    t.Log(AllResponse())
    time.Sleep(time.Second * 1)
    t.Log("After: ", runtime.NumGoroutine())

}

func TestPool(t *testing.T)  {
    pool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("create sync pool")
            return 10
        },
    }
    v1, _ := pool.Get().(int)
    fmt.Println(v1)
}


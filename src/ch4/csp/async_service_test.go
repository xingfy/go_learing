package async_service

import (
    "fmt"
    "testing"
    "time"
)

func service() string {
    time.Sleep(time.Millisecond * 50)
    return "Done"
}

func otherTask() {
    fmt.Println("working on something else")
    time.Sleep(time.Millisecond * 100)
    fmt.Println("task is done")
}

func TestService(t *testing.T) {
    fmt.Println(service())
    otherTask()
}

func AsyncService() chan string {
    //retCh := make(chan string)
    retCh := make(chan string, 1)

    go func() {
        ret := service()
        fmt.Println("return result")
        retCh <- ret
        fmt.Println("service exited")
    }()

    return retCh
}

func TestAsyncService(t *testing.T) {
    retCh := AsyncService()
    otherTask()
    fmt.Println(<-retCh)
}

func TestSelect(t *testing.T) {
    select {
    case ret := <-AsyncService():
        t.Log(ret)
    case <-time.After(time.Millisecond * 100):
        t.Error("time out")
    }

}

package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()

    for i := 0; i < 100000; i++ {
        go func(i int) {
            fmt.Println("用户第一,\n 拥抱变化,\n 协同增效,\n 正直诚信,\n 始终创业")
        }(i)
    }

    fmt.Println(time.Since(now))
}
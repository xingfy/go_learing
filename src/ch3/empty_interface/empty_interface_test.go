package empty_interface

import (
    "fmt"
    "testing"
)

func DoSomeThing(p interface{}) {
    //if i, ok := p.(int); ok {
    //    fmt.Println("int", i)
    //    return
    //}
    //if i, ok := p.(string); ok {
    //    fmt.Println("string", i)
    //    return
    //}
    //fmt.Println("unknown type")

    switch v := p.(type) {
    case int:
        fmt.Println("switch int", v)
    case string:
        fmt.Println("switch string", v)
    default:
        fmt.Println("switch unknown type")
    }

}

func TestEmptyInterface(t *testing.T) {
    DoSomeThing(1)
}

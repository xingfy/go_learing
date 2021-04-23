package ch5

import (
    "bytes"
    "testing"
)

func TestConcatStringByAdd(t *testing.T) {
    elems := []string{"1", "2", "3", "4", "5"}
    ret := ""
    for _, elem := range elems {
        ret += elem
    }
    t.Log(ret == "12345")
}

func TestConcatStringByBytesBuffer(t *testing.T) {
    elems := []string{"1", "2", "3", "4", "5"}
    var buf bytes.Buffer
    for _, elem := range elems {
        buf.WriteString(elem)
    }
    t.Log(buf.String() == "12345")
}

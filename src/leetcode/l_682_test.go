package leetcode

import (
    "fmt"
    "strconv"
    "testing"
)

func calPoints(ops []string) int {

    var res = make([]int, len(ops))
    var index = 0
    var result = 0

    for _, op := range ops {
        if op == "+" {
            res[index] = res[index-1] + res[index-2]
            index++
        } else if op == "D" {
            res[index] = res[index-1] * 2
            index++
        } else if op == "C" {
            res[index-1] = 0
            index--
        } else {
            number_int, _ := strconv.Atoi(op)
            res[index] = number_int
            index++
        }
        fmt.Println(op, res)
    }

    for _, r := range res {
        result += r
    }
    return result
}

func TestCalPoints(t *testing.T) {
    s := []string{"36", "28", "70", "65", "C", "+", "33", "-46", "84", "C"}
    calPoints(s)
}

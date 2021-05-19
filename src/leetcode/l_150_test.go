package leetcode

import (
    "fmt"
    "strconv"
    "testing"
)

type stack struct {
    num   [100]int
    index int
}

func (s *stack) Push(item int) {
    fmt.Println(s.index)
    s.num[s.index] = item
    s.index++
    fmt.Println(s.index)
}

func (s *stack) Pop() (int, bool) {
    if len(s.num) == 0 || s.index <= 0 {
        return 0, false
    } else {
        fmt.Println(s.num, s.index)
        s.index--
        item := s.num[s.index]
        fmt.Println(s.num, s.index)

        fmt.Println(s.num, s.index)
        return item, true
    }
}

func evalRPN(tokens []string) int {
    var helpNum stack

    for _, token := range tokens {
        val, err := strconv.Atoi(token)
        if err == nil {
            helpNum.Push(val)
            fmt.Println(helpNum)
        } else {
            num1, _ := helpNum.Pop()
            num2, _ := helpNum.Pop()

            switch token {
            case "+":
                helpNum.Push(num1 + num2)
            case "-":
                helpNum.Push(num2 - num1)
            case "*":
                helpNum.Push(num2 * num1)
            case "/":
                helpNum.Push(num2 / num1)
            }
        }
    }

    res, _ := helpNum.Pop()
    return res
}

func TestEvalRPN(t *testing.T) {
    tokens := []string{"2","1","+","3","*"}
    evalRPN(tokens)
}
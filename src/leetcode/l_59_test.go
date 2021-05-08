package leetcode

import (
    "testing"
)

func generateMatrix(n int) [][]int {
    var (
        res                      = make([][]int, n)
        left, right, top, bottom = 0, n - 1, 0, n - 1
        index                    = 0
    )

    for i := range res {
        res[i] = make([]int, n)
    }

    for left <= right && top <= bottom {
        for column := left; column <= right; column++ {
            index++
            res[top][column] = index
        }
        for row := top + 1; row <= bottom; row++ {
            index++
            res[row][right] = index
        }
        if left < right && top < bottom {

            for column := right - 1; column > left; column-- {
                index++
                res[bottom][column] = index
            }
            for row := bottom; row > top; row-- {
                index++
                res[row][left] = index
            }
        }
        left++
        right--
        top++
        bottom--
    }

    return res
}

func TestGenerateMatrix(t *testing.T) {
    res := generateMatrix(3)
    t.Log(res)
}
package leetcode

import (
    "fmt"
    "testing"
)

func rotate48(matrix [][]int) {

    n := len(matrix)
    if n == 0 {
        return
    }

    for i := 0; i < n/2; i++ {
        for j := 0; j < (n+1)/2; j++ {
            tmp := matrix[i][j]
            matrix[i][j] = matrix[n-1-j][i]
            matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
            matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
            matrix[j][n-1-i] = tmp
        }
    }

    fmt.Println(matrix)
}

func TestRotate48(t *testing.T) {
    var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    rotate48(matrix)
}

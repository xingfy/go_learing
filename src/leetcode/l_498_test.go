package leetcode

import "testing"

var index = 0

func findDiagonalOrder(mat [][]int) []int {

    index = 0

    var (
        // n为行数, m为列数
        endR, endC = len(mat) - 1, len(mat[0]) - 1
        res        = make([]int, (endR+1)*(endC+1))

        // false从左下到右上遍历
        fromUp = false

        tR, tC, dR, dC = 0, 0, 0, 0
    )

    if mat == nil || len(mat) == 0 {
        return res
    }

    for tR != endR+1 {
        printLevel(mat, tR, tC, dR, dC, fromUp, res)

        if tC == endC {
            tR = tR + 1
        } else {
            tC = tC + 1
        }

        if dR == endR {
            dC = dC + 1
        } else {
            dR = dR + 1
        }

        fromUp = !fromUp
    }

    return res
}

func printLevel(mat [][]int, tR int, tC int, dR int, dC int, fromUp bool, res []int) {
    if fromUp {
        for tR != dR+1 {
            res[index] = mat[tR][tC]
            index++
            tR++
            tC--
        }
    } else {
        for dR != tR-1 {
            res[index] = mat[dR][dC]
            index++
            dR--
            dC++
        }
    }
}

func TestFindDiagonalOrder(t *testing.T) {
    mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    diagonalOrder := findDiagonalOrder(mat)
    t.Log(diagonalOrder)
    diagonalOrder = findDiagonalOrder(mat)
}

/**
1 2 3
4 5 6
7 8 9
*/

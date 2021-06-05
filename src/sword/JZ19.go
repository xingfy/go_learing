package sword

func printMatrix(matrix [][]int) []int {
    // write code here

    var (
        rows, columns            = len(matrix), len(matrix[0])
        order                    = make([]int, rows*columns)
        index                    = 0
        left, right, top, bottom = 0, columns - 1, 0, rows - 1
    )

    for left <= right && top <= bottom {
        for column := left; column <= right; column++ {
            order[index] = matrix[top][column]
            index++
        }
        for row := top + 1; row <= bottom; row++ {
            order[index] = matrix[row][right]
            index++
        }

        if left < right && top < bottom {
            for column := right - 1; column > left; column-- {
                order[index] = matrix[bottom][column]
                index++
            }
            for row := bottom; row > top; row-- {
                order[index] = matrix[row][left]
                index++
            }
        }
        left++
        right--
        top++
        bottom--
    }

    return order
}

/**
1  2  3  4
5  6  7  8
9  10 11 12
13 14 15 16


1 2 3 4 8 12 16 15 14 14 9 5 6 7 11 10
*/

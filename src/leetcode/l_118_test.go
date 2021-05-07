package leetcode

func generate(numRows int) [][]int {

    ans := make([][]int, numRows)
    for idx := range ans {
        ans[idx] = make([]int, idx+1)
        ans[idx][0] = 1
        ans[idx][idx] = 1
        for j := 1; j < idx; j++ {
            ans[idx][j] = ans[idx-1][j] + ans[idx-1][j-1]
        }
    }
    return ans
}

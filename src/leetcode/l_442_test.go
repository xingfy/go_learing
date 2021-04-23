package leetcode

import (
    "fmt"
    "testing"
)

func findDuplicates(nums []int) []int {

    var res []int
    for idx, _ := range nums {
        num := nums[idx]
        if num < 0 {
            num = num * -1
        }

        if nums[num - 1] < 0 {
            res = append(res, num)
        } else {
            nums[num - 1] *= -1
        }
    }

    return res
}


func TestName(t *testing.T) {
    fmt.Println(100 ^ 100)
}

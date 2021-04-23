package leetcode

import (
    "sort"
    "testing"
)

func findErrorNums(nums []int) []int {

    var errNum = -1
    var repeatNum = 1

    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            errNum = nums[i]
        }
        if nums[i]-nums[i-1] > 1 {
            repeatNum = nums[i-1] + 1
        }
    }

    if nums[len(nums)-1] != len(nums) {
        repeatNum = len(nums)
    }

    return []int{errNum, repeatNum}
}

func TestFindErrorNums(t *testing.T) {

    var nums = []int{1, 3, 3}
    errorNums := findErrorNums(nums)
    t.Log(errorNums)
}

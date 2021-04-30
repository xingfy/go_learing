package leetcode

import "sort"

// TODO
func minMoves(nums []int) int {

    sort.Ints(nums)
    var count int = 0
    for idx, _ := range nums {
        count += nums[idx] - nums[0];
    }

    return count
}

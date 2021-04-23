package leetcode

import "testing"

func findMaxConsecutiveOnes(nums []int) int {

    var count = 0
    var maxCount = 0

    for _, num := range nums {
        if num == 1 {
            count++
            maxCount = If(maxCount > count, maxCount, count).(int)
        } else {
            count = 0
        }
    }

    return maxCount
}

func If(b bool, t, f interface{}) interface{} {
    if b {
        return t
    }
    return f
}

func TestMax(t *testing.T) {
    max := findMaxConsecutiveOnes([]int{1, 1, 0, 0, 1, 1, 1})
    t.Log(max)
}

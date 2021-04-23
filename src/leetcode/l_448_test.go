package leetcode

import (
    "fmt"
    "testing"
)

func findDisappearedNumbers(nums []int) (ans []int) {

    n := len(nums)
    for _, v := range nums {
        v = (v - 1) % n
        fmt.Printf("v:%d \t", v)
        nums[v] += n
    }
    for i, v := range nums {
        if v < n {
            ans = append(ans, i+1)
        }
    }
    return
}

func findDisappearedNumbersV2(nums []int) (ans []int) {

    for _, v := range nums {
        if v-1 < 0 {
            v = v * -1
        }
        if nums[v-1] > 0 {
            nums[v-1] *= -1
        }

    }

    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            ans = append(ans, i+1)
        }
    }
    return
}

func TestFindDisappearedNumbers(t *testing.T) {

    nums := []int{2, 1, 3, 3, 4, 6}
    fmt.Println(nums)
    ans := findDisappearedNumbers(nums)
    t.Log(ans)
}

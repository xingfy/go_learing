package leetcode

/**
前缀 乘 后缀
*/
func productExceptSelf(nums []int) []int {

    var (
        n = len(nums)

        // 前缀和数组
        before = make([]int, n)

        // 后缀和数组
        after = make([]int, n)

        res = make([]int, n)
    )

    before[0] = 1
    after[n-1] = 1

    for i := 1; i < len(nums); i++ {
        before[i] = before[i-1] * nums[i-1]
    }

    for i := len(nums) - 2; i >= 0; i-- {
        after[i] = after[i+1] * nums[i+1]
    }

    for idx := range nums {
        res[idx] = before[idx] * after[idx]
    }

    return res
}

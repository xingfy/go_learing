package leetcode

func rotate(nums []int, k int) {
    n := len(nums)
    if n == 1 {
        return
    }

    // 定义辅助数组
    copyNums := make([]int, n)

    for idx, _ := range nums {
        copyNums[(idx+k)%n] = nums[idx]
    }

    copy(nums, copyNums)
}

func rotateV2(nums []int, k int) {
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(a []int) {
    n := len(a)

    for i := 0; i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

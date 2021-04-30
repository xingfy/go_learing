package leetcode

func firstMissingPositive(nums []int) int {

    n := len(nums)
    if n == 0 {
        return 0
    }

    for idx, num := range nums {
        if num <= 0 {
            nums[idx] = n + 1
        }
    }

    for idx, _ := range nums {
        num := abs(nums[idx])
        if num <= n {
            nums[num - 1] = -abs(nums[num - 1])
        }
    }

    for idx, _ := range nums {
        if nums[idx] > 0 {
            return idx + 1
        }
    }

    return n + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

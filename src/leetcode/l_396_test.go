package leetcode

import "math"

func maxRotateFunction(nums []int) int {

    maxRotate := math.MinInt32
    n := len(nums) - 1
    for n >= 0 {
        maxRotateReverse(nums)
        maxRotateReverse(nums[:1])
        maxRotateReverse(nums[1:])

        tmp := 0
        for idx, num := range nums {
            tmp += idx * num
        }

        if maxRotate < tmp {
            maxRotate = tmp
        }

        n--
    }

    return maxRotate
}

func maxRotateReverse(a []int) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}


func maxRotateFunctionV2(nums []int) int {

    sum, rSum, n := 0, 0, len(nums)

    for idx, num := range nums {
        sum += num
        rSum += idx * num
    }

    ans := rSum
    cur := n - 1
    for cur > 0 {
        rSum += sum - nums[cur] * n
        if ans < rSum {
            ans = rSum
        }

        cur--
    }

    return ans
}
package leetcode

func checkPossibility(nums []int) bool {

    if len(nums) == 1 {
        return true
    }

    // 判断初始条件
    sign := nums[0] <= nums[1]

    for i := 1; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            if sign {
                if nums[i+1] >= nums[i-1] {
                    nums[i] = nums[i+1]
                } else {
                    nums[i+1] = nums[i]
                }
                sign = true
            } else {
                return false
            }
        }
    }

    return true
}

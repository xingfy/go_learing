package leetcode

/**
快慢指针
*/
func moveZeroes(nums []int) {

    left := 0
    right := 0
    n := len(nums)

    for right < n {
        if nums[right] != 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
        right++
    }
}

/**

0 1 0 3 12 0 4 0
1 0 0 3 12 0 4 0
1 0 3 0 12 0 4 0
1 3 0 0 12 0 4 0
1 3 0 12 0 0 4 0
1 3 12 0 0 0 4 0
1 3 12 0 0 4 0 0
1 3 12 0 4 0 0 0
1 3 12 4 0 0 0 0


0 1 0 3 12 0 4 0
1 0 0 3 12 0 4 0
1 3 0 0 12 0 4 0
1 3 12 0 0 0 4 0
1 3 12 4 0 0 0 0

*/

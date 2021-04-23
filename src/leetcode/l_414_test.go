package leetcode

import "math"

func thirdMax(nums []int) int {

    var maxNum =  math.MinInt64
    var secNum =  math.MinInt64
    var thirdNum =  math.MinInt64

    for _, num := range nums {
        if num == maxNum || num == secNum || num == thirdNum {
            continue
        }

        if num > maxNum {
            thirdNum = secNum
            secNum = maxNum
            maxNum = num
        } else if num > secNum {
            thirdNum = secNum
            secNum = num
        } else if num > thirdNum {
            thirdNum = num
        }
    }

    if (thirdNum ==  math.MinInt64) {
        return maxNum
    } else {
        return thirdNum
    }
}
package leetcode

type entry struct {
    count int
    left  int
    right int
}

func findShortestSubArray(nums []int) int {

    mp := map[int]entry{}
    for index, num := range nums {
        if v, has := mp[num]; has {
            v.count++
            v.right = index
            mp[num] = v
        } else {
            mp[num] = entry{1, index, index}
        }
    }

    maxCnt := 0
    minLen := 0
    for _, v := range mp {
        if v.count > maxCnt {
            maxCnt = v.count
            minLen = v.right - v.left + 1
        } else if v.count == maxCnt {
            if v.right - v.left + 1 < minLen {
                minLen = v.right - v.left + 1
            }
        }
    }

    return minLen
}

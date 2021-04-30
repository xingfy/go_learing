package leetcode

import "sort"

func hIndex(citations []int) int {

    sort.Ints(citations)

    res := 0
    l := len(citations) - 1
    for res < len(citations) && citations[l] > res {
        res++
        l--
    }

    return res ;
}

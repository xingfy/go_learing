package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {

    sort.Slice(people, func(i, j int) bool {
        a, b := people[i], people[j]
        if a[0] != b[0] {
            return a[0] > b[0]
        } else {
            return a[1] < b[1]
        }
    })

    var ans [][]int

    for _, p := range people {
        idx := p[1]
        ans = append(ans[:idx], append([][]int{p}, ans[idx:]...)...)
    }

    return ans
}

// [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// [7,0], [7,1], [6,1], [5,0], [5,2], [4,4]


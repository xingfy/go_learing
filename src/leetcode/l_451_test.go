package leetcode

func frequencySort(s string) string {

    var count = make([]int, 26)

    for i := range s {
        count[s[i]-'a']++
    }
    
    return ""
}

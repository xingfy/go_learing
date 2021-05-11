package leetcode

func countSegments(s string) int {

    if len(s) == 0 {
        return 0
    }

    count := 0

    for i := 1; i < len(s); i++ {
        if s[i-1] != ' ' && s[i] == ' ' {
            count++
        }
    }
    if s[len(s)-1] != ' ' {
        count++
    }

    return count
}

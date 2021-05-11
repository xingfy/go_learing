package leetcode

import "strings"

func isPalindrome(s string) bool {
    s = strings.ToLower(s)

    var (
        left  = 0
        right = len(s) - 1
    )

    for left < right {
        for left < right && !isalnum(s[left]) {
            left++
        }
        for left < right && !isalnum(s[right]) {
            right--
        }

        if left < right {
            if s[left] == s[right] {
                left++
                right--
            } else {
                return false
            }
        }
    }
    return true
}

func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

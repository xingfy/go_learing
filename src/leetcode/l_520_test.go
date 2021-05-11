package leetcode

func detectCapitalUse(word string) bool {


    var (
        // 大写字母标记
        upperWord = false
        // 小写字母标记
        lowerWord = false
    )

    // 反向遍历
    for i := len(word) - 1; i >= 0; i-- {
        // 如果该字母是大写
        if word[i] >= 'A' && word[i] <= 'Z' {
            // 如果之前出现过大写和小写, 则不正确
            if upperWord && lowerWord {
                return false
            }
            upperWord = true
        } else {
            if upperWord {
                return false
            }
            lowerWord = true
        }
    }
    return true
}

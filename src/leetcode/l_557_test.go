package leetcode

func reverseWords(s string) string {

    str := []byte(s)
    var i = 0
    var j = 0
    var isStart = true
    for start := 0; start < len(s); {

        if isStart {
            i = start
            isStart = false
        }
        if s[start] == ' ' {
            j = start - 1

            for i < j {
                str[i], str[j] = str[j], str[i]
            }

            isStart = true
        }
        start++
    }

    return string(str)
}

package leetcode

func reverseStr(s string, k int) string {

    str := []byte(s)
    for start := 0; start < len(s); start += 2 * k {
        var i = start
        j := start + k - 1
        if j > len(s) - 1 {
            j = len(s) - 1
        }

        for i < j {
            str[i], str[j] = s[j], s[i]
            i++
            j--
        }
    }

    return string(str)
}
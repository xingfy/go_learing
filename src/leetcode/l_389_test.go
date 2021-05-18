package leetcode

func findTheDifference(s string, t string) byte {

    cnt := [26]int{}
    var res byte

    for i := range s {
        cnt[s[i]-'a']++
    }

    for _, v := range t {
        if cnt[v-'a'] == 0 {
            res = byte(v)
            return res
        }
        cnt[v-'a']--
    }

    return res
}

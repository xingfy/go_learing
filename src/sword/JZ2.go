package sword

func replaceSpace( s string ) string {
    // write code here
    l := len(s)
    if l == 0 {
        return ""
    }
    var help string

    for i := 0; i < l; i++ {
        if s[i] == ' ' {
            help += "%20"
        } else {
            help += string(s[i])
        }
    }
    return help
}

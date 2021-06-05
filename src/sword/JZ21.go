package sword

import "fmt"

func IsPopOrder(pushV []int, popV []int) bool {
    // write code here
    if len(pushV) != len(popV) {
        return false
    }

    var (
        help = make([]int, len(pushV))
        idx  = 0
        i, j = 0, 0
    )

    for i < len(pushV) {
        if pushV[i] != popV[j] {
            help[idx] = pushV[i]
            i++
            idx++

            fmt.Println("push", help, i, j, idx)
        } else {
            i++
            j++
            idx--
            fmt.Println(help[idx], popV[j])
            for idx >= 0 && help[idx] == popV[j] {
                idx--
                j++

                fmt.Println("pop", help, i, j, idx)
            }
            idx++
        }
    }

    return idx == 0
}

func main() {
    res := IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
    fmt.Println(res)
}

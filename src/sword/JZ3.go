package sword

type ListNode struct {
    Val  int
    Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
    // write code here
    var res = make([]int, 100)
    var index int

    for head != nil {
        cur := head.Val
        res[index] = cur
        index++
        head = head.Next
    }

    return res
}

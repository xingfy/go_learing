package sword

func ReverseList(pHead *ListNode) *ListNode {
    // write code here
    var help *ListNode
    cur := pHead

    for cur != nil {
        // 记录next节点
        next := cur.Next
        cur.Next = help
        help = cur
        cur = next
    }

    return help
}

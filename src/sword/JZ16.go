package sword

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
    // write code here
    var head = &ListNode{
        Val: 0,
        Next: nil,
    }
    var res = head

    cur1, cur2 := pHead1, pHead2

    for cur1 != nil && cur2 != nil {
        if cur1.Val > cur2.Val {
            res.Next = cur2
            cur2 = cur2.Next
        } else {
            res.Next = cur1
            cur1 = cur1.Next
        }

        res = res.Next
    }

    if cur1 != nil {
        res.Next = cur1
    } else if cur2 != nil {
        res.Next = cur2
    }

    return head.Next
}

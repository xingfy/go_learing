package sword

func FindKthToTail(pHead *ListNode, k int) *ListNode {
    // write code here

    fast, slow := pHead, pHead

    for i := 0; i < k; i++ {
        fast = fast.Next
    }

    for fast != nil {
        fast, slow = fast.Next, slow.Next
    }

    return slow
}

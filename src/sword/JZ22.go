package sword

func PrintFromTopToBottom( root *TreeNode ) []int {
    // write code here
    if root == nil {
        return []int{}
    }

    var (
        queue = make([]*TreeNode, 0)
        res = make([]int, 0)
    )
    queue = append(queue, root)

    for len(queue) > 0 {
        // 弹出第一个元素
        node := queue[0]
        queue = queue[1:]

        res = append(res, node.Val)

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    return res
}
package sword

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
    // write code here

    if len(pre) == 0 || len(vin) == 0 {
        return nil
    }

    root := &TreeNode{pre[0], nil, nil}
    i := 0

    for ; i < len(vin); i++ {
        if vin[i] == pre[0] {
            break
        }
    }
    root.Left = reConstructBinaryTree(pre[1:len(vin[:i])+1], vin[:i])
    root.Right = reConstructBinaryTree(pre[len(vin[:i])+1:], vin[i+1:])

    return root
}

package sword

func Mirror(pRoot *TreeNode) *TreeNode {
    // write code here

    if pRoot == nil {
        return nil
    }

    pRoot.Left, pRoot.Right = Mirror(pRoot.Right), Mirror(pRoot.Left)

    return pRoot
}

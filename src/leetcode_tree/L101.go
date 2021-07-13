package leetcode_tree

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isSymmetric2(root, root)
}

func isSymmetric2(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }

    if left.Val != right.Val {
        return false
    }

    return isSymmetric2(left.Left, right.Right) && isSymmetric2(left.Right, right.Left)
}

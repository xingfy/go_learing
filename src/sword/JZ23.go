package sword

/**
判断输入的数组是否是某二叉搜索树的后续遍历的结果
*/
func VerifySquenceOfBST(sequence []int) bool {
    // write code here

    return recur(sequence, 0, len(sequence)-1)
}

func recur(seq []int, i int, j int) bool {
    if i >= j {
        return true
    }
    p := i
    for seq[p] < seq[j] {
        p++
    }
    m := p
    for seq[p] > seq[j] {
        p++
    }

    return p == j && recur(seq, i, m-1) && recur(seq, m, j - 1)
}


/**
[4,8,6,12,16,14,10]

     10
  6      14
4   8  12  16
*/

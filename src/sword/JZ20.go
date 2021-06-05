package sword

var (
    stack = make([]int, 0)
    min   = make([]int, 0)
)

/**
每次min中只压入当前长度栈的最小的元素¬
 */
func Push(node int) {
    // write code here
    if len(min) > 0 && min[len(min)-1] < node {
        min = append(min, min[len(min)-1])
    } else {
        min = append(min, node)
    }
    stack = append(stack, node)
}
func Pop() {
    // write code here
    stack = stack[:len(stack) - 1]
    min = min[:len(min) - 1]
}
func Top() int {
    // write code here
    return stack[len(stack) - 1]
}
func Min() int {
    // write code here
    return min[len(min) - 1]
}

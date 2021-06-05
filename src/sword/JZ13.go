package sword

func reOrderArray(array []int) []int {
    // write code here

    var (
        res = make([]int, len(array))
        index int = 0
    )

    for _, num := range array {

        if num%2 == 1 {
            res[index] = num
            index++
        }
    }
    for _, num := range array {
        if num%2 == 0 {
            res[index] = num
            index++
        }
    }

    return res
}

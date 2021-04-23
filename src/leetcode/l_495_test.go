package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 0 || duration <= 0 {
        return 0
    }

    var cur = 0

    // 每次求上个timer的持续时间
    for index := 0; index < len(timeSeries) - 1; index++ {
        if timeSeries[index+1] - timeSeries[index] < duration {
            cur = cur + timeSeries[index+1] - timeSeries[index]
        } else {
            cur = cur + duration
        }
    }

    return cur + duration
}

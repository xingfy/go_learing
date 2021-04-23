package err_test

import (
    "errors"
    "testing"
)

func GetFib(n int) ([]int, error) {
    if n <= 0 {
        return nil, errors.New("param illegal")
    }

    fibList := []int{1, 1}
    for i := 2; i < n; i++ {
        fibList = append(fibList, fibList[i-1]+fibList[i-2])
    }
    return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
    if v, err := GetFib(-10); err != nil {
        t.Log(err)
    } else {
        t.Log(v)
    }

}

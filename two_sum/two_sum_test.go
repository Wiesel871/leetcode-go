package two_sum

import (
    "testing"
)

func TestTwoSum1(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    if result[0] != 0 || result[1] != 1 {
        t.Error("TwoSum failed. Expected [0, 1], Got ", result)
    }
}

func TestTwoSum2(t *testing.T) {
    nums := []int{3, 2, 4}
    target := 6
    result := twoSum(nums, target)
    if result[0] != 1 || result[1] != 2 {
        t.Error("TwoSum failed. Expected [1, 2], Got ", result)
    }
}

func TestTwoSum3(t *testing.T) {
    nums := []int{3, 3}
    target := 6
    result := twoSum(nums, target)
    if result[0] != 0 || result[1] != 1 {
        t.Error("TwoSum failed. Expected [0, 1], Got ", result)
    }
}

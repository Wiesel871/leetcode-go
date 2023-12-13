package two_sum

func twoSum(nums []int, target int) []int {
    numl := len(nums)
    if numl == 2 {
        return []int{0, 1}
    }
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{0, 0}
}

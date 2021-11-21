func sortedSquares(nums []int) []int {
    for i, v := range nums {
        nums[i] *= v
    }
    sort.Ints(nums)
    return nums
}
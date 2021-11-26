func searchInsert(nums []int, target int) int {
    // Brute force
    for i, v := range nums {
        if v == target || v > target {
            return i
        }
    }
    return len(nums)
}
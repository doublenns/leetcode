func removeElement(nums []int, val int) int {
    overwriteIndex := 0
    for _, v := range nums {
        if v != val {
            nums[overwriteIndex] = v
            overwriteIndex++
        }
    }
    return overwriteIndex
}
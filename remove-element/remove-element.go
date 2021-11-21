func removeElement(nums []int, val int) int {
    j := len(nums) - 1
    for i := 0; i < len(nums) && i <= j ; i++ {
        if nums[i] == val {
            nums[i], nums[j] = nums[j], nums[i]
            i--
            j--
        }
    }
    return j+1
}
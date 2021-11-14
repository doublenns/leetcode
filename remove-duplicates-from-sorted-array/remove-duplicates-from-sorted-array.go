func removeDuplicates(nums []int) int {
    l := len(nums)
    if l <= 1 {
        return l
    }
    
    i := 0
    for j := 1; j < l; j++ {
        if nums[i] != nums[j] {
            nums[i+1] = nums[j]
            i++
        }
    }
    return i+1
}
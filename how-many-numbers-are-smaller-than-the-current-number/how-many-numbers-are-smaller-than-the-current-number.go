func smallerNumbersThanCurrent(nums []int) []int {
    // Brute force implementation
    result := make([]int, len(nums))
    for i, v := range nums{
        counter := 0
        for _, compare := range nums {
            if compare < v {
                counter++
            }
        }
        result[i] = counter
    }
    return result
}
func missingNumber(nums []int) int {
    // Initialize sum w/ length of input slice to account for the
    // last element in a slice w/ no missing numbers
    sum, numsSum := len(nums), 0
    for i := 0; i <len(nums); i++ {
        sum += i
        numsSum += nums[i]
    }
    return sum - numsSum
}
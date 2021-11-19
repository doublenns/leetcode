func largestOddNumber(num string) string {
    nums := strings.Split(num, "")

    for i := len(nums) -1 ; i >= 0; i-- {
        n, _ := strconv.Atoi(nums[i])
        if n % 2 != 0 {
            return strings.Join(nums[:i+1], "")
        }
    }
    
    return ""
}
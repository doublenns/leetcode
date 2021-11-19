func largestOddNumber(num string) string {
    nums := strings.Split(num, "")
    last := -1
    for i := len(nums) -1 ; i >= 0; i-- {
        n, _ := strconv.Atoi(nums[i])
        if n % 2 != 0 {
            last = i
            break
        }
    }
    return strings.Join(nums[:last+1], "")
}
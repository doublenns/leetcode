func largestOddNumber(num string) string {
    for i := len(num) -1 ; i >= 0; i-- {
        n, _ := strconv.Atoi(string(num[i]))
        if n % 2 != 0 {
            return string(num[:i+1])
        }
    }
    
    return ""
}
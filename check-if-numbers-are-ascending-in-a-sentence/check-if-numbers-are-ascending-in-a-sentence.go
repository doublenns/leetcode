func areNumbersAscending(s string) bool {
    compare := 0
    for _, v := range strings.Fields(s) {
        num, err := strconv.Atoi(v)
        if err == nil {
            if num <= compare {
                return false
            }
            compare = num
        }
    }
    return true
}
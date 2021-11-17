func sumSlice(sl []string) int {
    result := 0
    for _, v := range sl {
        num, _ := strconv.Atoi(v)
        result += num
    }
    return result
}

func addDigits(num int) int {
    str := strconv.Itoa(num)
    if len(str) != 1 {
        nums := strings.Split(str, "")
        newNum := sumSlice(nums)
        return addDigits(newNum)
    }
    result, _ := strconv.Atoi(str)
    return result
}
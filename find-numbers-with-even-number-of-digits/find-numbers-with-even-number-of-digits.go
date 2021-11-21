func numberOfDigits(n int) int {
    count := 1
    for n >= 10 {
        n /= 10
        count++
    }
    return count
}

func findNumbers(nums []int) int {
    var count int
    for _, v := range nums {
        if numberOfDigits(v) % 2 == 0{
            count++
        }
    }
    return count
}
func sumOfUnique(nums []int) int {
    seen := make(map[int]int)
    
    for _, v := range nums {
        seen[v]++
    }
    
    var result int
    for i, v := range seen {
        if v == 1 {
            result += i
        }
    }
    return result
}
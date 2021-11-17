func findDisappearedNumbers(nums []int) []int {
    // Brute force solution
    seen := make(map[int]struct{})
    exists := struct{}{}
    
    for _, v := range nums{
        seen[v] = exists
    }
    
    var result []int
    for i := 1; i <= len(nums); i++ {
        if _, ok := seen[i]; !ok {
            result = append(result, i)
        }
    }
    return result
}
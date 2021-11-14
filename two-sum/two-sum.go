func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    var result []int
    

	for i, v := range nums {
        if t, ok := seen[target - v]; ok {
            result = append(result, t, i)
            break
        } else {
            seen[v] = i
        }
	}
	return result
}
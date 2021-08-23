func twoSum(nums []int, target int) []int {
	var result []int
	counter := 0
	for i, v := range nums {
		counter++
		for i2, v2 := range nums[(i + 1):] {
			if v+v2 == target {
				result = append(result, i, i2+counter)
				return result
			}
		}
	}
	return result
}
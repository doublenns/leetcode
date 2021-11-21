func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if len(fmt.Sprintf("%d", n))%2 == 0 {
			count++
		}
	}
	return count
}
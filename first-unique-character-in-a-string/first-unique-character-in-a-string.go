func firstUniqChar(s string) int {
    // Solution using a closure func
	slice_to_map := func(s string) map[rune]int {
		frequency := make(map[rune]int)
		for _, v := range s {
            frequency[v]++
		}
		return frequency
	}

	m := slice_to_map(s)
	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
    return -1
}
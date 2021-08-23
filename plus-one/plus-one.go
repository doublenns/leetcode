func plusOne(digits []int) []int {
	var plusOne func(int)
	plusOne = func(i int) {
		if digits[i] == 9 {
			digits[i] = 0
			if i > 0 {
				plusOne(i - 1)
			} else {
				digits = append(digits, 0)
				copy(digits[1:], digits)
				digits[0] = 1
			}
		} else {
			digits[i]++
		}
	}

	plusOne(len(digits) - 1)
    return digits
}
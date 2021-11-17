func areNumbersAscending(s string) bool {
	compare := 0

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			var number int
			number, i = readNumber(s, i)
			if number <= compare {
				return false
			}

			compare = number
		}
	}

	return true
}

func readNumber(s string, start int) (int, int) {
	result := 0
	for i := start; i < len(s); i++ {
		if !isDigit(s[i]) {
			return result, i
		}

		result = result*10 + int(s[i]-'0')
	}

	return result, len(s)
}

func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}
func fizzBuzz(n int) []string {
	// Solutin using a switch
    var result []string
	for i := 1; i <= n; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
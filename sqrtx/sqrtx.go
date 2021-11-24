const DELTA = 0.01

// Function abs returns the absolute value of a number passed into it
func abs(num float64) float64 {
	if num < 0 {
		num = 0 - num
	}
	return num
}

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    
    var a float64 = float64(x) / 2
    var temp float64
    for abs(a - temp) > DELTA {
        temp = a
        a = a - ((a*a) - float64(x)) / (2*a)
    }
    return int(a)
}
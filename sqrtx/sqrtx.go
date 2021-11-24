const DELTA = 0.01

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    
    var a float64 = float64(x) / 2
    var temp float64
    for math.Abs(a - temp) > DELTA {
        temp = a
        a = a - ((a*a) - float64(x)) / (2*a)
    }
    return int(a)
}
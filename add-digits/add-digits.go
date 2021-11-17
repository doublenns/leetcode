// https://math.stackexchange.com/questions/2102877/explanation-of-digital-root-sum-formula
// O(1) time/space solution
func addDigits(num int) int {
    return 1 + ( (num - 1) % 9 )
}
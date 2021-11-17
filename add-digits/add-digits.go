// Recursive Solution

// func sumSlice(sl []string) int {
//     result := 0
//     for _, v := range sl {
//         num, _ := strconv.Atoi(v)
//         result += num
//     }
//     return result
// }

// func addDigits(num int) int {
//     str := strconv.Itoa(num)
//     if len(str) != 1 {
//         nums := strings.Split(str, "")
//         newNum := sumSlice(nums)
//         return addDigits(newNum)
//     }
//     result, _ := strconv.Atoi(str)
//     return result
// }


// https://math.stackexchange.com/questions/2102877/explanation-of-digital-root-sum-formula
// O(1) time/space solution
func addDigits(num int) int {
    return 1 + ( (num - 1) % 9 )   
}
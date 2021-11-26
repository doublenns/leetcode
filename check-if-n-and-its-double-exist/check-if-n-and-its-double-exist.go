func checkIfExist(arr []int) bool {
    // O(n^2) solution (nested for loop)
    
    for i, v := range arr {
        for j, x := range arr {
            if i != j && x % 2 == 0 && v * 2 == x {
                return true
            }
        }
    }

    return false
}
func checkIfExist(arr []int) bool {
    for i, v := range arr {
        if v % 2 == 0 {
            for j, x := range arr {
                if i != j && x * 2 == v {
                    return true
                }
            }
        }
    }
    return false
}
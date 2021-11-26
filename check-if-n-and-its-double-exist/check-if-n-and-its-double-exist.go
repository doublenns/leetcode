func checkIfExist(arr []int) bool {
    seen := make(map[float64]struct{})
    exists := struct{}{}
    
    for _, v := range arr {
        if _, ok := seen[float64(v)*2.0]; ok {
            return true
        } else if _, ok := seen[float64(v)/2.0]; ok {
            return true
        } else {
            seen[float64(v)] = exists
        }
    }
    return false
}
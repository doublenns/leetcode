func checkIfExist(arr []int) bool {
    seen := make(map[int]struct{})
    exists := struct{}{}
    
    for _, v := range arr {
        if _, ok := seen[v*2]; ok {
            return true
        } else if v % 2 == 0 {
            if _, ok := seen[v/2]; ok {
                return true
            }
        }
        seen[v] = exists
    }
    return false
}
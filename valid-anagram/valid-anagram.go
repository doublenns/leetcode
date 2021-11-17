func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    check1 := make(map[rune]int)
    check2 := make(map[rune]int)
    
    for _, v := range s {
        check1[v]++
    }
    for _, v := range t {
        check2[v]++
    }
    
    for i, v := range check1 {
        if check2[i] != v {
            return false
        }
    }
    
    return true
}
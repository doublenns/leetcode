func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    check := make(map[rune]int)
    
    for _, v := range s {
        check[v]++
    }
    for _, v := range t {
        check[v]--
    }
    
    for _, v := range check {
        if v != 0 {
            return false
        }
    }
    
    return true
}
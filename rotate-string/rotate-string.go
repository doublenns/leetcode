func isSubstring(s, t string) bool {
    for i := 0; i < len(t); i++ {
        for j := 0; j < len(s) && i+j < len(t) && s[j] == t[i + j]; j++{
            if j == len(s) - 1{
                return true
            } 
        }
    }
    return false
}

func rotateString(s string, goal string) bool {
    if len(goal) != len(s) {
        return false
    }
    check := s + s
    return isSubstring(goal, check)
}
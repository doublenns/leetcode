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

func numOfStrings(patterns []string, word string) int {
    count := 0
    for _, s := range patterns {
        if isSubstring(s, word) {
            count++
        }
    }
    return count
}
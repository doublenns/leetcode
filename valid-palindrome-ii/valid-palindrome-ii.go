func validPalindrome(s string) bool {
    var validSubPalindrome func(int, int, bool) bool
    validSubPalindrome = func(start, end int, skipAvailable bool) bool {
        for i, j := start, end; i <= j; i, j = i+1, j-1 {
            if s[i] != s[j] {
                if skipAvailable {
                    return validSubPalindrome(i+1, j, false) || validSubPalindrome(i, j-1, false)
                } else {
                    return false
                }   
            }
        }
        return true
    }
    
    return validSubPalindrome(0, len(s)-1, true)
}
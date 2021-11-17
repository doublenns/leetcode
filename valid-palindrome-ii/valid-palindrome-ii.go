func validPalindrome(s string) bool {
    var isPalindrome func(int, int, bool) bool
    isPalindrome = func(start, end int, skipAvailable bool) bool {
        for i, j := start, end; i <= j; i, j = i+1, j-1 {
            if s[i] != s[j] {
                if skipAvailable {
                    return isPalindrome(i+1, j, false) || isPalindrome(i, j-1, false)
                }
                return false
            }
        }
        return true
    }
    
    return isPalindrome(0, len(s)-1, true)
}
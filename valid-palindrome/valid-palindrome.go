// The "strings" import in previous submission slowed down execution considerably

func removeNonAlphanumerics(s string) string {
    diff := 'A' - 'a'
    
    var result strings.Builder
    for _, r := range s {
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
            result.WriteRune(r)
        } else if r >= 'A' && r <= 'Z' {
            result.WriteRune(r - diff)
        }
    }
    
    // Using bytes, where possible, instead of runes to see if faster
    // for i := 0; i < len(s); i++ {
    //     b := s[i]
    //     if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')  {
    //         result.WriteByte(b)
    //     } else if b >= 'A' && b <= 'Z' {
    //         result.WriteRune(rune(b) - diff)
    //     }
    // }
    
    return result.String()
}

func isPalindrome(s string) bool {
    p := removeNonAlphanumerics(s)
    for i, j := 0, len(p)-1; i <= j; i, j = i+1, j-1 {
        if p[i] != p[j] {
            return false
        }
    }
    return true
}
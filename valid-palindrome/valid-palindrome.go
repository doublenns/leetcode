import "strings"

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
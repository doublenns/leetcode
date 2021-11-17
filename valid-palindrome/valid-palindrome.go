func isPalindrome(s string) bool {
    diff := 'A' - 'a'
    j := len(s)-1;
    
    for i, r := range s {
        if i >= j {
            break
        }
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
            if r >= 'A' && r <= 'Z' {
                r -= diff
            }
        } else {
            continue
        }
                                                            
        
        for j >= i {
            c := rune(s[j])
            if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
                if c != r {
                    return false
                }
            } else if c >= 'A' && c <= 'Z' {
                c -= diff
                if c != r {
                    return false
                }
            }   else {
                j--
                continue
            }
            j--
            break
        }
        
    }
    return true
}
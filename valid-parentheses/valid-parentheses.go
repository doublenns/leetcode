func isValid(s string) bool {
    braces := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }
    
    var stack []rune
    for _,c := range s{
        if _, ok := braces[c]; ok {
            stack = append(stack, c)
        } else {
            if len(stack) <= 0 {
                return false
            } else {
                l := len(stack) -1
                t := stack[l]
                if braces[t] != c {
                    return false
                }
                stack = stack[:l]
            }
        }
    }
    
    return len(stack) == 0
}
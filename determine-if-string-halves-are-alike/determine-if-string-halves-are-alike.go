func halvesAreAlike(s string) bool {
    exists := struct{}{}
    vowels := map[byte]struct{}{
        'a': exists, 'A': exists,
        'e': exists, 'E': exists,
        'i': exists, 'I': exists,
        'o': exists, 'O': exists,
        'u': exists, 'U': exists,
    }
    
    firstHalf, secondHalf := 0, 0
    
    for i, j := 0, len(s) -1; i < j ; i, j = i+1, j-1 {
        if _, ok := vowels[s[i]]; ok {
            firstHalf++
        }
        if _, ok := vowels[s[j]]; ok {
            secondHalf++
        }
    }
    return firstHalf == secondHalf
}
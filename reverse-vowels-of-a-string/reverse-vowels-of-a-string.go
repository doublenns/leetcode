func reverseVowels(s string) string {
    exists := struct{}{}
    vowels := map[rune]struct{}{
        'a': exists, 'A': exists,
        'e': exists, 'E': exists,
        'i': exists, 'I': exists,
        'o': exists, 'O': exists,
        'u': exists, 'U': exists,
    }
    var vowelIndexes []int
    
    for i, r := range s {
        if _, ok := vowels[r]; ok {
            vowelIndexes = append(vowelIndexes, i)
        }
    }
    sb := []byte(s)
    for i, j := 0, len(vowelIndexes) -1; i<j; i, j = i+1, j-1 {
        sb[vowelIndexes[i]], sb[vowelIndexes[j]] = sb[vowelIndexes[j]], sb[vowelIndexes[i]]
    }
    
    return string(sb)
}
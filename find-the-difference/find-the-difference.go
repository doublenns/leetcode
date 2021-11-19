func findTheDifference(s string, t string) byte {
    if len(t) == 1 {
        return t[0]
    }
    
    compare := make(map[rune]int)
    
    for _, r := range t {
        compare[r]++
    }
    
    for _, r := range s {
        compare[r]--
        if compare[r] == 0 {
            delete(compare, r)
        }
    }
    fmt.Println(compare)
    var result byte
    for k, _ := range compare {
        result = string(k)[0]
    }
    
    return result
}
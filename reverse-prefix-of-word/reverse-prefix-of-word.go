func reversePrefix(word string, ch byte) string {
    var end int
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            end = i
            break
        }
    }
    
    bs := []byte(word)
    
    for i, j := 0, end; i < j; i, j = i+1, j-1 {
        bs[i], bs[j] = bs[j], bs[i]
    }
    return string(bs)
}
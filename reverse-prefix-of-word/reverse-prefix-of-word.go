func reversePrefix(word string, ch byte) string {
    // Get first index of ch in word
    end := strings.IndexByte(word, ch)
    // Strings are immutable, so convert word to slice of bytes
    bs := []byte(word)
    
    // Reverse the chars in word up to index of ch
    for i, j := 0, end; i < j; i, j = i+1, j-1 {
        bs[i], bs[j] = bs[j], bs[i]
    }
    
    // Convert slice of bytes back into a string and return
    return string(bs)
}
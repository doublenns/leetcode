func reverseWords(s string) string {
    words := strings.Fields(s)
    var result []string
    for _, word := range words {
        w := []byte(word)
        for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
            w[i], w[j] = w[j], w[i]
        }
        result = append(result, string(w))
        
    }
    return strings.Join(result, " ")
}
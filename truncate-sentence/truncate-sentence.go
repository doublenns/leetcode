func truncateSentence(s string, k int) string {
    words := strings.Fields(s)
    return strings.Join(words[:k], " ")
}
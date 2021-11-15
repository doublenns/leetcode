func checkIfPangram(sentence string) bool {
    // Create a set of all runes in sentence and return whether 26 different
    // characters used. (Constraint indicates no punctuation or whitespace)
    check := make(map[rune]struct{})
    var exists = struct{}{}
    
    for _, r := range sentence {
        check[r] = exists
    }
    return len(check) == 26
}
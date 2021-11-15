func checkIfPangram(sentence string) bool {
    // Create a set of all runes in sentence and return whether 26 different
    // characters used. (Constraint indicates no punctuation or whitespace)
    seen := make(map[rune]struct{})
    var exists = struct{}{}
    
    for _, r := range sentence {
        seen[r] = exists
    }
    return len(seen) == 26
}
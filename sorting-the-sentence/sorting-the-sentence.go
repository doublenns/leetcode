import "strings"

func sortSentence(s string) string {

    words := strings.Fields(s)
    // Preallocate space in slice for each word of sentence
    result := make([]string, len(words))

    for _, word := range words {
        length := len(word)
        posChar := word[length - 1]
        // '0' (rune["0"]) == 48; '1' == 49
        // Subtract 1 at end to have the pos 0 indexed
        pos := posChar - '0' - 1
        result[pos] = word[:length-1]
    }
    
    // Return the slice of words, separated by a space
    return strings.Join(result, " ")
}
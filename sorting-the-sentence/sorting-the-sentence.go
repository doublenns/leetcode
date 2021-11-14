import "strings"

// Create a type of struct to store the "naked" word and the int that indicates
// the word's position in the original un-shuffled sentence
type word struct {
    str string
    pos rune
}

func separateWords(s string) []word {
    // Store the word structs in a slice
    var words []word
    
    for _, v := range strings.Fields(s) {
        length := len(v)
        w := word{
            pos: []rune(v)[length -1],
            str: string([]rune(v)[:length-1]),
        }
        words = append(words, w)
    }
    
    return words
}

func sortSentence(s string) string {

    words := separateWords(s)

    // Sort the slice of structs based on the "pos" attribute
    sort.Slice(words, func(i, j int) bool {
		return words[i].pos < words[j].pos
	})
    
    // Iterate through the slice of structs and place them in a new slice, in 
    // correct sentence order
    var result []string
    for _, st := range words {
        result = append(result, st.str)
    }
    // Return the slice of words, separated by a space
    return strings.Join(result, " ")
}
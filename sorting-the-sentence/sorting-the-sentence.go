import "strings"

func sortSentence(s string) string {
    type word struct {
        str string
        pos rune
    }
    var words []word
    
    for _, v := range strings.Fields(s) {
        length := len(v)
        w := word{
            pos: []rune(v)[length -1],
            str: string([]rune(v)[:length-1]),
        }
        words = append(words, w)
    }

    sort.Slice(words, func(i, j int) bool {
		return words[i].pos < words[j].pos
	})
    
//     var result strings.Builder
//     for _, st := range words {
//         result.WriteString(st.str)
//     }
    
//     return result.String()
    
    var result []string
    for _, st := range words {
        result = append(result, st.str)
    }
    
    return strings.Join(result, " ")
}
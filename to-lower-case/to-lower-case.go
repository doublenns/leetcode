import "strings"

func toLowerCase(s string) string {
    // 'a' == 97; 'A' == 65
    diff := 'a' -'A'
    var result strings.Builder
    
    for _, r := range s {
        if ('A' <= r && r <= 'Z') {
			result.WriteRune(r + diff)
        } else {
            result.WriteRune(r)
        }
	}
	return result.String()
}
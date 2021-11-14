import "strings"

func lengthOfLastWord(s string) int {
    w := strings.Fields(s)
    return len(w[len(w)-1])
}
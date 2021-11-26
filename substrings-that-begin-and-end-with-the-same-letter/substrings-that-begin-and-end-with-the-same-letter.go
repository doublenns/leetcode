func numberOfSubstrings(s string) int64 {
    seen := make(map[rune]int64)
    var result int64
    for _, v := range s {
        seen[v]++
        result += seen[v]
    }
    return result
}
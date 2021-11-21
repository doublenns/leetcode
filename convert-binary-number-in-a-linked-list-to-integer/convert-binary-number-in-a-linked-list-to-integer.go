func getDecimalValue(head *ListNode) int {
    // Add all digits to a string
    var str strings.Builder
    for curr := head; curr != nil; curr = curr.Next {
        if curr.Val == 1 {
            str.WriteString("1")
        } else if curr.Val == 0 {
            str.WriteString("0")
        }
    }
    
    // Convert base 2 string rep of int to base 10 int
    result, _ := strconv.ParseInt(str.String(), 2, 32)
    return int(result)
}
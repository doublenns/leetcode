func hasCycle(head *ListNode) bool {
    // Hashmap lookup solution
    seen := make(map[*ListNode]struct{})
    exists := struct{}{}
    
    for curr := head; curr != nil; curr = curr.Next {
        if _, ok := seen[curr]; ok {
            return true
        }
        seen[curr] = exists
    }
    return false
}
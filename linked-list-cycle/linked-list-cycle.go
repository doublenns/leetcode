/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]struct{})
    exists := struct{}{}
    
    curr := head
    for curr != nil && curr.Next != nil {
        if _, ok := seen[curr]; ok {
            return true
        }
        seen[curr] = exists
        curr = curr.Next
    }
    return false
}
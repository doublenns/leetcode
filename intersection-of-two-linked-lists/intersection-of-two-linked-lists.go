/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    seen := make(map[*ListNode]struct{})
    exists := struct{}{}
    
    for curr := headA; curr !=nil; curr = curr.Next {
        seen[curr] = exists
    }
    for curr := headB; curr !=nil; curr = curr.Next {
        if _, ok := seen[curr]; ok {
            return curr
        }
    }
    return nil
}
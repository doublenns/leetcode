/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    curr := head
    for curr.Next != nil {
        temp := curr.Next
        for temp.Val == curr.Val {
            if temp.Next != nil {
                temp = temp.Next
            } else {
                curr.Next = nil
                return head
            }
        }
		curr.Next = temp
		curr = temp
    }
    return head
}
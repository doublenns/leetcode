/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    for curr := head; curr != nil; curr = curr.Next {
        for curr.Next != nil && curr.Next.Val == val {
            curr.Next = curr.Next.Next
        }
    }
    return head
}
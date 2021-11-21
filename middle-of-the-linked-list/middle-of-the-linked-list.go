/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var tail int
    for curr := head; curr != nil; curr = curr.Next {
        tail++
    }
    var temp *ListNode
    for curr, tail := head, (tail/2)+1; curr != nil && tail > 0; curr, tail = curr.Next, tail-1 {
        temp = curr
    }
    return temp
}
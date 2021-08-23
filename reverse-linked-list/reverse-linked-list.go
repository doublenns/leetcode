/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type LinkedList struct {
	head   *ListNode
}

func (l *LinkedList) Prepend(n *ListNode) {
	if l.head == nil {
		n.Next = nil
        l.head = n
	} else {
		n.Next = l.head
		l.head = n
	}
}

func reverseList(head *ListNode) *ListNode {
    l := LinkedList{}
    for head != nil {
        temp := head.Next
        l.Prepend(head)
        head = temp
    }
    return l.head
}
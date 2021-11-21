/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    
    var res *ListNode
    if l1.Val < l2.Val {
        res = l1
        l1 = l1.Next
    } else {
        res = l2
        l2 = l2.Next
    }
    
    temp := res
    for p1, p2 := l1,l2 ; p1 !=nil || p2 !=nil; {
        if p1 == nil {
            temp.Next = p2
            temp = temp.Next
            p2 = p2.Next
        } else if p2 == nil {
            temp.Next = p1
            temp = temp.Next
            p1 = p1.Next
        } else if p1.Val <= p2.Val {
            temp.Next = p1
            temp = temp.Next
            p1 = p1.Next
        } else {
            temp.Next = p2
            temp = temp.Next
            p2 = p2.Next
        }
    }
    return res
}
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    for curr := head; curr != nil; {
        for curr.Next != nil && curr.Next.Val == val {
            curr.Next = curr.Next.Next
            continue
        }
        curr = curr.Next
    }
    return head
}
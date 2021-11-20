func deleteDuplicates(head *ListNode) *ListNode {
    curr := head
    for curr != nil && curr.Next != nil {
        temp := curr.Next
        if temp.Val == curr.Val {
            curr.Next = temp.Next
        } else {
            curr = temp
        }
    }
    return head
}


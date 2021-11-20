func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    curr := head
    for curr.Next != nil {
        temp := curr.Next
        if temp.Val == curr.Val {
            curr.Next = temp.Next
        } else {
            curr = temp
        }
    }
    return head
}


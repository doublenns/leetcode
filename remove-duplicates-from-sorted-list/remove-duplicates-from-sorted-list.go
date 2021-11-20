func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
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
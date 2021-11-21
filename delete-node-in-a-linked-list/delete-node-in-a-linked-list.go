func deleteNode(node *ListNode) {
    d := node.Next
    node.Val, node.Next = node.Next.Val, node.Next.Next
    d.Next = nil
}
func printLinkedListInReverse(head ImmutableListNode) {
    for curr := head; curr != nil; curr = curr.getNext() {
        defer curr.printValue()
    }

}
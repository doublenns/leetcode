/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */
func printLinkedListInReverse(head ImmutableListNode) {
    var st []ImmutableListNode
    
    for curr := head; curr != nil; curr = curr.getNext() {
        st = append(st, curr)
    }
    for i := len(st) -1 ; i >= 0 ; i-- {
        st[i].printValue()
    }
}
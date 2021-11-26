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
type Stack struct {
	data []ImmutableListNode
}

func (st *Stack) Push(node ImmutableListNode) bool {
	st.data = append(st.data, node)
	return true
}

func (st *Stack) Pop() (ImmutableListNode, error) {
	length := len(st.data)
	if length == 0 {
		return nil, errors.New("empty stack")
	}
	rv := st.data[length-1]
	st.data = st.data[:length-1]
	return rv, nil
}




func printLinkedListInReverse(head ImmutableListNode) {
    var st []ImmutableListNode
    
    for curr := head; curr != nil; curr = curr.getNext() {
        st = append(st, curr)
    }
    for i := len(st) -1 ; i >= 0 ; i-- {
        st[i].printValue()
    }
    
}
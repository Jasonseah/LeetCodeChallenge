package list_node

/**
 * for illustrate purpose this will be on top of the function usually
 * should be outside from this class
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * construct a node object from a list of array
 */
func ConstructNode(nodes []int) ListNode {
	var l = ListNode{}
	l.Val = nodes[0]

	var p = &l
	for i := 1; i < len(nodes); i++ {
		p.Insert(nodes[i])
		p = p.Next
	}

	return l
}

/**
 *
 */
func (l *ListNode) Insert(value int) *ListNode {
	l.Next = &ListNode{Val: value}

	return l
}

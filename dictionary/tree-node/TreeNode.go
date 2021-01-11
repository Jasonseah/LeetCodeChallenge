package tree_node

/**
 * for illustrate purpose this will be on top of the function usually
 * should be outside from this class
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * construct a node object from a list of array
 */
func ConstructNode(nodes []int) TreeNode {
	var t = TreeNode{}
	t.Val = nodes[0]

	for i := 1; i < len(nodes); i++ {
		t.Insert(nodes[i])
	}

	return t
}

/**
 * recursively create node
 * warning I only able to create 3 layer at most
 *      1
 *   2    3
 *  4 5  6 7
 *  further than this you will be expecting everything will be added to the left
 *  reason being at the line 55 I check if left of the left and right have value it will switch to right
 *  but also line 64 I check if the right of the left and right have value it will switch to left
 *  with this case you will see the line 64 will be always true base on the diagram
 *  hence if always true it will always insert to node left
 *  unless I check t.left.left.left and t.right.right.right and so on
 */
func (t *TreeNode) Insert(value int) error {

	if t.Left == nil {
		t.Left = &TreeNode{Val: value}
		return nil
	}

	if t.Right == nil {
		t.Right = &TreeNode{Val: value}
		return nil
	}

	if t.Left != nil {
		if t.Left.Left != nil && t.Left.Right != nil {
			return t.Right.Insert(value)
		}

		return t.Left.Insert(value)
	}

	if t.Right != nil {
		if t.Right.Left != nil && t.Right.Right != nil {
			return t.Left.Insert(value)
		}

		return t.Right.Insert(value)
	}

	return nil
}

/**
 *
 */
//func (t *TreeNode) Flatten() []int {
//
//
//}

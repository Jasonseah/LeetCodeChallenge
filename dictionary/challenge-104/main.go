package challenge_104

import (
	. "stream-project.com/jasonseah/leet-code-challenge/dictionary/tree-node"
)

/**
 * run time 4 ms + ~4.5mb
 */
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftBaseRoot := 1
	leftBaseRoot += MaxDepth(root.Left)

	rightBaseRoot := 1
	rightBaseRoot += MaxDepth(root.Left)

	if leftBaseRoot > rightBaseRoot {
		return leftBaseRoot
	} else {
		return rightBaseRoot
	}
}


/**
 * run time 4 ms + ~4.5mb
 * similar approach
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1+max(maxDepth(root.Right), maxDepth(root.Left))
}


func max(right, left int)int{
	if right > left {
		return right
	}

	return left
}

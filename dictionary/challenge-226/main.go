package challenge_226

import (
	. "stream-project.com/jasonseah/leet-code-challenge/dictionary/tree-node"
)

/**
 * run time 4 ms + ~4.5mb
 * LOL Leet Code literally say this is faster than 100.00% of Go online submissions (YAY !)
 */
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	oriLeft := root.Left
	oriRight := root.Right

	root.Left = oriRight
	root.Right = oriLeft

	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)

	return root
}

package challenge_617

import (
	. "stream-project.com/jasonseah/leet-code-challenge/dictionary/tree-node"
)

/**
 * run time 56 ms + ~8mb
 * lesson learned too much if else basically if we remove the 1st if sentence
 * to make the similar if else with the advance attempt then we will get nearly the same ms
 */
func MergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil || t2 == nil {
		if t1 != nil {
			return t1
		}

		if t2 != nil {
			return t2
		}

		if t1 == nil && t2 == nil {
			return nil
		}
	}

	t1.Val += t2.Val

	if t1.Left != nil || t2.Left != nil {
		if t1.Left == nil {
			t1.Left = &TreeNode{Val: 0}
		}

		if t2.Left == nil {
			t2.Left = &TreeNode{Val: 0}
		}

		t1.Left = MergeTree(t1.Left, t2.Left)
	}

	if t1.Right != nil || t2.Right != nil {
		if t1.Right == nil {
			t1.Right = &TreeNode{Val: 0}
		}

		if t2.Right == nil {
			t2.Right = &TreeNode{Val: 0}
		}

		t1.Right = MergeTree(t1.Right, t2.Right)
	}

	return t1
}

/**
 * run time 28 ms + ~8mb
 */
func AdvanceMergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	t1.Left = AdvanceMergeTrees(t1.Left, t2.Left)
	t1.Right = AdvanceMergeTrees(t1.Right, t2.Right)

	return t1
}

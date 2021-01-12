package challenge_206

import (
	. "stream-project.com/jasonseah/leet-code-challenge/dictionary/list-node"
)

/**
 * my brute force xD
 * run time 36ms & ~20mb
 */
func ReverseList(head *ListNode) *ListNode {

	var reverseArray []int
	root := head

	for root != nil {
		reverseArray = append([]int{root.Val}, reverseArray...)
		root = root.Next
	}

	if len(reverseArray) < 1 {
		return nil
	}

	var l = ListNode{}
	l.Val = reverseArray[0]

	loopHead := &l
	for i := 1; i < len(reverseArray); i++ {
		loopHead.Next = &ListNode{Val: reverseArray[i]}
		loopHead = loopHead.Next
	}

	return &l
}


/**
 * understand the concept saving current and prev and continue to pass on to the next object
 * run time 4ms & ~3mb
 */
func AdvanceReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

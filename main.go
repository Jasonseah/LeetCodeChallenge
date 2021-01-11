package main

import (
	"fmt"
	challenge_104 "stream-project.com/jasonseah/leet-code-challenge/dictionary/challenge-104"
	tree_node "stream-project.com/jasonseah/leet-code-challenge/dictionary/tree-node"
)

func main() {
	// var array = []int{1, 2, 3, 4}
	// fmt.Println(challenge_1480.AdvanceVersion(array))

	//var array = [][]int{{2, 8}, {8, 8}, {7, 4}}
	//var challenge = challenge_1476.Constructor(array)

	//fmt.Println(challenge.GetValue(1, 0))
	//challenge.UpdateSubrectangle(1, 1, 1, 1, 4)
	//fmt.Println(challenge.GetArray())

	//var array = []int{2, 3, 5, 1, 3}
	//fmt.Println(challenge_1431.KidsWithCandies(array, 3))

	//var array = []int{1, 2, 3, 1, 1, 3}
	//fmt.Println(challenge_1512.AdvanceNumIdenticalPairs(array))

	//var array = []int{1, 2, 3, 1, 1, 3}
	//fmt.Println(challenge1470.Shuffle(array, 3))

	//fmt.Println(challenge_1108.DefangIPaddr("127.0.0.1"))

	//fmt.Println(challenge_771.NumJewelsInStonesAdvance("zA", "aaAAbbbA"))

	//fmt.Println(challenge_1342.NumberOfSteps(14))

	//var array = []int{4, 5, 6, 7, 0, 2, 1, 3}
	//fmt.Println(challenge_1528.RestoreString("codeleet", array))

	//fmt.Println(challenge_1281.SubtractProductAndSum(234))

	//var array = []int{1, 2, 3, 4}
	//fmt.Println(challenge_1313.DecompressRLElist(array))

	//var array = []int{8, 1, 2, 2, 3}
	//fmt.Println(challenge_1365.SmallerNumbersThanCurrent(array))

	//t1 := tree_node.ConstructNode([]int{1, 2, 3, 4, 5})
	//t2 := tree_node.ConstructNode([]int{8, 3, 0, 1, 2, 3})
	//
	//challenge_617.MergeTree(&t1, &t2)
	//fmt.Println(t1.Flatten())

	//3,9,20,null,null,15,7
	root := tree_node.ConstructNode([]int{3, 9, 20, 0, 0, 15, 7})
	//t2 := tree_node.ConstructNode([]int{8, 3, 0, 1, 2, 3})
	//
	fmt.Println(challenge_104.MaxDepth(&root))
	//fmt.Println(t1.Flatten())
}

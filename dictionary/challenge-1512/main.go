package challenge_1512

/**
 * brute force which is my 1st attempt
 */
func NumIdenticalPairs(nums []int) int {
	var array [][]int = nil

	for key, value := range nums {
		for i := key + 1; i < len(nums); i++ {
			if value == nums[i] {
				var pair = []int{key, i}
				array = append(array, pair)
			}
		}
	}

	return len(array)
}

/**
 * []int{1, 2, 3, 1, 1, 3}
 * algorithm way
 * 0 += 0 [0,1]
 * 0 += 0 [0,1,1]
 * 0 += 0 [0,1,1,1]
 * 0 += 1 [0,2,1,1] // 1
 * 1 += 2 [0,3,1,1] // 3
 * 3 += 1 [0,3,1,2] // 4
 */
func AdvanceNumIdenticalPairs(nums []int) int {
	count, tmp := 0, make([]int, 101)

	// []int{1, 2, 3, 1, 1, 3}
	for _, v := range nums {
		count += tmp[v]
		tmp[v]++
	}

	return count
}

package challenge_1470

/**
 * normal way
 */
func Shuffle(nums []int, n int) []int {
	array := make([]int, len(nums))

	keyToAppend := -1
	for i := 0; i < n; i++ {
		keyToAppend += 1
		array[keyToAppend] = nums[i]
		keyToAppend += 1
		array[keyToAppend] = nums[i+n]
	}

	return array
}

/**
 * var array = []int{1, 2, 3, 1, 1, 3, 4, 3}
 * 1231 1343
 * 11 231 343
 * 11 23 31 43
 * 11 23 34 13
 */
func AdvanceShuffle(nums []int, n int) []int {
	count := 0
	for i := 1; i < len(nums)-1; i++ {
		j := i
		for nums[i] >= 0 {
			if j < n {
				j = j * 2
			} else {
				j = (j-n)*2 + 1
			}
			nums[i], nums[j] = nums[j], -nums[i]
			count++
			if count == len(nums)-2 {
				goto end
			}
		}
	}
end:
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = -nums[i]
	}
	return nums
}

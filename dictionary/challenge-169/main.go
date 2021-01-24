package challenge_169

/**
 * time complexity = O(n)
 * space complexity = 4n + 4n + 16
 * run time 20 ms + ~6mb
 * similar approach
 */
// nums = 4n
func majorityElement(nums []int) int {
	// encounterCount = 4n
	encounterCount := make(map[int]int)
	// both are fixed if 4 bytes
	var maxKey = 0
	var maxValue = 0

	// num 4 bytes
	for _, num := range nums {
		encounterCount[num] = encounterCount[num] + 1
	}

	// key, count each 4 bytes = 8 bytes
	for key, count := range encounterCount {
		if count > maxValue {
			maxValue = count
			maxKey = key
		}
	}

	return maxKey
}

/**
 * time complexity = O(n)
 * space complexity = 4n + 16
 * run time 16 ms + ~6mb
 * similar approach
 */
// nums = 4n
func MajorityElement(nums []int) int {
	// 4 bytes
	candidate := nums[0]
	// 4 bytes
	count := 1

	// 8 bytes
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count++
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

package challenge_1313

/**
 * 1st attempt
 */
func DecompressRLElist(nums []int) []int {
	var output []int

	loop := 0
	for loop < len(nums) {

		if nums[loop] >= 1 {
			freq := 0
			for freq < nums[loop] {
				output = append(output, nums[loop+1])
				freq++
			}
		}

		loop += 2
	}

	return output
}

/**
 * low memory usage
 */
func decompressRLElist(nums []int) []int {
	resultLen := 0
	//O(n/2)
	for i, count := range nums {
		if i%2 != 0 {
			continue
		}
		resultLen = resultLen + count
	}

	// O(n)
	result := make([]int, resultLen)

	pointer := 0
	// O(n/2)
	for i, count := range nums {
		if i%2 != 0 {
			continue
		}
		num := nums[i+1]
		for j := 0; j < count; j++ {
			result[pointer] = num
			pointer++
		}
	}

	return result
}

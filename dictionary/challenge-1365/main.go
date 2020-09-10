package challenge_1365

import "sort"

// O(N2) - N * N
func SmallerNumbersThanCurrent(nums []int) []int {
	var comparedResults []int

	for key, value := range nums {
		comparedCount := 0

		for keyToCompare, valueToCompare := range nums {
			if key != keyToCompare && value > valueToCompare {
				comparedCount++
			}
		}

		comparedResults = append(comparedResults, comparedCount)
	}

	return comparedResults
}

// O(N)
func AdvanceSmallerNumbersThanCurrent(nums []int) []int {
	counter := make([]int, 101)
	for i := range nums {
		counter[nums[i]] += 1
	}
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i - 1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = 0
			continue
		}
		nums[i] = counter[nums[i]-1]
	}
	return nums
}

// N * log(N)
func smallerNumbersThanCurrent2(nums []int) []int {
	type pair struct {
		index int
		value int
	}
	arr := make([]pair, len(nums))
	for i, val := range nums {
		arr[i] = pair{i, val}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].value < arr[j].value
	})

	dups := 0
	lastNum := -1
	for i, pair:= range arr {
		if pair.value == lastNum {
			dups += 1
		} else {
			dups = 0
		}
		nums[pair.index] = i - dups
		lastNum = pair.value
	}
	return nums
}

// N * log(N) with binary search
func smallerNumbersThanCurrent3(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	// fmt.Println(sorted)
	// fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] = sort.SearchInts(sorted, nums[i])
	}
	return nums
}

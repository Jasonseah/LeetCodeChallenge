package challenge_1480

/**
a solution attempt by me
 */
func NormalVersion(input []int) []int {
	//arr := make([]int, len(input)) // [0,0,0,0]

	var arr []int = nil

	for key, value := range input {
		if key == 0 {
			arr = append(arr, value)
		} else {
			arr = append(arr, arr[key-1]+value)
		}
	}

	return arr
}

/**
a solution provided by
https://github.com/qiyuangong/leetcode/blob/master/java/1480_Running_Sum_of_1d_Array.java
 */
func AdvanceVersion(input []int) []int {
	//arr := make([]int, len(input)) // [0,0,0,0]
	if len(input) <= 1 {
		return input
	}

	for key := range input {
		if key >= 1 {
			input[key] += input[key-1]
		}
	}

	return input
}

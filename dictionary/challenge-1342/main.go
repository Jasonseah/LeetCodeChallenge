package challenge_1342

/**
 *
 */
func NumberOfSteps(num int) int {
	count := 0

	if num <= 2 {
		return num
	}

	//fmt.Println(recursiveReduceNo(num))
	for num != 2 {
		num = recursiveReduceNo(num)
		count++
	}

	return count + 2
}

func recursiveReduceNo(num int) int {
	if num%2 == 0 {
		return num / 2
	} else {
		return num - 1
	}
}


/**
 * advance method
 */
func AdvanceNumberOfSteps (num int) int {
	if num <= 1 {
		return num
	}

	if num % 2 == 0 {
		return 1 + AdvanceNumberOfSteps(num / 2)
	}

	return 1 + AdvanceNumberOfSteps(num - 1)
}

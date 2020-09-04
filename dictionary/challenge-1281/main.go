package challenge_1281

import (
	"strconv"
)

/**
 * 1st attempt brute force
 */
func SubtractProductAndSum(n int) int {
	stringOfInt := strconv.Itoa(n)
	multiplyNo := 0
	additionalNo := 0

	for key, val := range stringOfInt {
		int, _ := strconv.Atoi(string(val))

		if key == 0 {
			multiplyNo = int
		} else {
			multiplyNo *= int
		}

		additionalNo += int
	}

	return multiplyNo - additionalNo
}

/**
 * advance approch
 */
func AdvanceSubtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 {
		product *= n % 10
		sum += n % 10
		n /= 10
	}
	return product - sum
}

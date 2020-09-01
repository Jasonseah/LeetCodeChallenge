package challenge_771

import (
	"fmt"
	"strings"
)

/**
 * 1st attempt brute force
 */
func NumJewelsInStones(J string, S string) int {
	count := 0

	for _, value := range S {
		for _, jewel := range J {
			if strings.Compare(string(value), string(jewel)) == 0 {
				count++
			}
		}
	}

	return count
}

/**
 * more smarter way to do it
 * please look at the challenge 1521 had the nearly same approach
 */
func NumJewelsInStonesAdvance(J string, S string) int {
	jewels := make([]bool, 256)
	for i := 0; i < len(J); i++ {
		jewels[J[i]] = true
		fmt.Println(J[i])
	}

	fmt.Println(jewels)

	count := 0

	for i := 0; i < len(S); i++ {
		if jewels[S[i]] {
			count++
		}
	}
	return count

}

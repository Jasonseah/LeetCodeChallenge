package challenge_1528

import "strings"

/**
 * 1st attempt
 */
func RestoreString(s string, indices []int) string {
	array := make([]string, len(indices))

	for key, val := range s {
		array[indices[key]] = string(val)
	}

	return strings.Join(array[:], ",")
}

/**
 * advance attempt
 */
func AdvanceRestoreString(s string, indices []int) string {
	array := make([]byte, len(indices))

	for key, position := range indices {
	array[position] = s[key]
	}

	return string(array)
}

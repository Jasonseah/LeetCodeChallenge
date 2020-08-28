package challenge_1108

import (
	"fmt"
	"strings"
)

/**
 * brute force which is my 1st attempt
 */
func DefangIPaddr(address string) string {
	var modifiedAddressBytes []byte
	addressBytes := []byte(address)

	for _, alphabetByte := range addressBytes {
		if alphabetByte == 46 {
			modifiedSeperater := []byte("[.]")
			modifiedAddressBytes = append(modifiedAddressBytes, modifiedSeperater...)
		} else {
			modifiedAddressBytes = append(modifiedAddressBytes, alphabetByte)
		}
	}

	return string(modifiedAddressBytes)
}

/**
 * native string replace
 */
func DefangIPaddrNative(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

/**
 * an advance way to string modify
 */
func DefangIPaddrAdvance(address string) string {
	s := strings.FieldsFunc(address, func(c rune) bool {
		return c == '.'
	})

	return fmt.Sprintf("%s[.]%s[.]%s[.]%s", s[0], s[1], s[2], s[3])
}

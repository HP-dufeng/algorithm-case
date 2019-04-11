package reversestring

import (
	"fmt"
	"strings"
)

// ReverseByForloop reverse string just use for loop
func ReverseByForloop(str string) string {
	var result string
	for _, c := range str {
		result = fmt.Sprintf("%c%s", c, result)
	}

	return result
}

// ReverseBySlice : reverse string by slice,
// First convert string to slice,
// and simply loop through the first half of the array,
// swapping each element in turn with its mirror counterpart
func ReverseBySlice(str string) string {
	if len(str) <= 0 {
		return str
	}

	result := strings.Split(str, "")
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return strings.Join(result, "")
}

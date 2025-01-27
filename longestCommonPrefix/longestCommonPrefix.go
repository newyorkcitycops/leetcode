package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	// Base case: empty string array
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefixRecursive(strs, 0, len(strs)-1)
}

// Comparing prefixes from both halves
func commonPrefix(left, right string) string {
	minLength := min(len(left), len(right))
	for i := range minLength {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLength]
}

// Splitting array into two halves (last and current)
func longestCommonPrefixRecursive(strs []string, left, right int) string {
	if left == right {
		return strs[left]
	}

	// Mid is essential for halves comparison
	mid := (left + right) / 2
	lcpLeft := longestCommonPrefixRecursive(strs, left, mid)
	lcpRight := longestCommonPrefixRecursive(strs, mid+1, right)
	return commonPrefix(lcpLeft, lcpRight)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flo", "flower", "flop"}))
}

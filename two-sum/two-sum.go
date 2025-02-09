package main

import (
	"fmt"
	"strconv"
	"strings"
)

func twoSum(
	sequence []int,
	target int,
) []string {
	var combinations []string

	for i := 0; i < len(sequence)-1; i++ {
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i]+sequence[j] == target {
				combinations = append(
					combinations,
					strconv.Itoa(i),
					strconv.Itoa(j),
				)
				break
			}
		}
	}

	return combinations
}

func main() {
	sequence := []int{1, 2, 3, 4, 5, 6}

	combinations := twoSum(sequence, 3)

	fmt.Printf("[%s]\n", strings.Join(combinations, ", "))
}

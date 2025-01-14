package main

import (
	"fmt"
	"math"
)

func palindrome(number int) bool {
	reverse := 0
	numberTemp := number

	for number != 0 {
		remainder := number % 10
		reverse = reverse*10 + remainder
		number /= 10
	}

	return int(math.Abs(float64(reverse))) == numberTemp
}

func main() {
	fmt.Println(palindrome(-121))
}

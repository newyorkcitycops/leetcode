package main

import "fmt"

func isValid(s string) bool {
	l := make([]rune, 0, len(s))
	m := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, e := range s {
		if len(l) == 0 || l[len(l)-1] != m[e] {
			l = append(l, e)
			continue
		}

		l = l[:len(l)-1]
	}

	return len(l) == 0
}

func main() {
	fmt.Println(isValid("([])"))
}

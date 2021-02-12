package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("[[[)))"))
}

func isValid1(s string) bool {
	mp := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0, len(s))
	for _, b := range s {
		if mp[b] != 0 {
			if len(stack) == 0 || stack[len(stack)-1] != mp[b] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}

func isValid(s string) bool {
	for {
		l := len(s)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)

		if l == len(s) {
			return false
		}

		if len(s) == 0 {
			return true
		}
	}
}

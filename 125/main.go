package main

import (
	"fmt"
)

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z'
}

func equal(c1, c2 byte) bool {
	if c1 < 'a' {
		c1 += 'a' - 'A'
	}
	if c2 < 'a' {
		c2 += 'a' - 'A'
	}
	return c1 == c2
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !isAlpha(s[i]) {
			i++
		} else if !isAlpha(s[j]) {
			j--
		} else {
			if !equal(s[i], s[j]) {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

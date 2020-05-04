package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	//fmt.Printf("s, p: %s, %s\n", s, p)
	if len(p) == 0 {
		return len(s) == 0
	}

	subs := strings.Split(p, "*")
	//fmt.Println(subs)
	if len(s) < len(subs[0]) {
		return false
	}
	for i := 0; i < len(subs[0]); i++ {
		if subs[0][i] != '?' && subs[0][i] != s[i] {
			return false
		}
	}
	if len(subs) == 1 {
		return len(s)==len(subs[0])
	}

	si := len(subs[0])
	for k := 1; k < len(subs)-1; k++ {
		subStr := subs[k]
		if len(s[si:]) < len(subStr) {
			return false
		}
		var i, j int
		for i = si; i <= len(s)-len(subStr); i++ {
			for j = 0; j < len(subStr); j++ {
				if subStr[j] != '?' && subStr[j] != s[i+j] {
					break
				}
			}
			if j == len(subStr) {
				break
			}
		}
		if i > len(s)-len(subStr) {
			return false
		}
		si = i + len(subStr)
	}


	if len(s[si:]) < len(subs[len(subs)-1]) {
		return false
	}
	for i := 0; i < len(subs[len(subs)-1]); i++ {
		subStr := subs[len(subs)-1]
		if subStr[len(subStr)-1-i] != '?' && subStr[len(subStr)-1-i]!= s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("ho", "ho**"))
	fmt.Println(isMatch("abefcdgiescdfimde", "ab*cd?i*de"))
	fmt.Println(isMatch("aaaa", "***a"))
}

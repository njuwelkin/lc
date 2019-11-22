package main

import "fmt"

func c(n int) int {
	ret := 1
	for n > 0 {
		ret <<= 1
		n--
	}
	return ret - 1
}

func findSubstringInWraproundString(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := make([]int, 26)

	var i, j int
	for i, j = 0, 1; j < len(s); j++ {
		if s[j] != (s[j-1]+1-'a')%26+'a' {
			for ; i < j; i++ {
				l := j - i
				if l > count[s[i]-'a'] {
					count[s[i]-'a'] = l
				}
			}
		}
	}
	for ; i < j; i++ {
		l := j - i
		if l > count[s[i]-'a'] {
			count[s[i]-'a'] = l
		}
	}

	ret := 0
	for i = 0; i < 26; i++ {
		ret += count[i]
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findSubstringInWraproundString("a"))
	fmt.Println(findSubstringInWraproundString("cac"))
	fmt.Println(findSubstringInWraproundString("cabc"))
}

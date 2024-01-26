package main

import "fmt"

func letterCasePermutation(s string) []string {
	ret := []string{}
	buf := make([]byte, len(s))
	var f func(int)
	f = func(depth int) {
		if depth == len(s) {
			ret = append(ret, string(buf))
			return
		}
		c := s[depth]
		buf[depth] = c
		f(depth + 1)
		if c <= 'Z' && c >= 'A' {
			buf[depth] = c - 'A' + 'a'
			f(depth + 1)
		} else if c <= 'z' && c >= 'a' {
			buf[depth] = c - 'a' + 'A'
			f(depth + 1)
		}
	}
	f(0)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

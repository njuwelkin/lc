package main

import "fmt"

func findSubstringInternal(s string, words []string) []int {
	n := len(words)
	if len(s) == 0 || n == 0 {
		return []int{}
	}
	l := len(words[0])
	ret := []int{}
	for i := 0; i < l; i++ {
		if len(s)-i < l*n {
			break
		}
		m := map[string]int{}
		for _, w := range words {
			m[w]++
		}
		str := s[i:]
		for j, k := 0, 0; k <= len(str)-n; {
		}
	}
}

func main() {
	fmt.Println("vim-go")
}

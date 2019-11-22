package main

import "fmt"

func minWindow(s string, t string) string {
	count := make([]int, 256)
	countDiff := 0
	for i := range t {
		if count[t[i]] == 0 {
			countDiff++
		}
		count[t[i]]--
	}
	contains := 0
	retI, retLen := 0, 1<<62
	var i, j int
	for i, j = 0, 0; j < len(s); {
		if contains < countDiff {
			count[s[j]]++
			if count[s[j]] == 0 {
				contains++
			}
			j++
		} else {
			if retLen > (j - i) {
				retI, retLen = i, j-i
			}
			count[s[i]]--
			if count[s[i]] < 0 {
				contains--
			}
			i++

		}
	}
	for contains >= countDiff {
		if retLen > (j - i) {
			retI, retLen = i, j-i
		}
		count[s[i]]--
		if count[s[i]] < 0 {
			contains--
		}
		i++
	}
	if retLen > len(s) {
		return ""
	}
	return s[retI : retI+retLen]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
}

package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	count := make([]int, 26)
	minCount := len(s)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 && count[i] < minCount {
			minCount = count[i]
		}
	}
	if minCount >= k {
		return len(s)
	}

	ret := 0
	for i := 0; i <= len(s)-k; {
		if count[s[i]-'a'] < k {
			i++
			continue
		}
		var j int
		for j = i; j < len(s) && count[s[j]-'a'] >= k; j++ {
		}
		if j-i > ret {
			ret = max(longestSubstring(s[i:j], k), ret)
		}
		i = j
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}

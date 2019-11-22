package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	visited := make([]bool, 256)
	ret := 0
	for i, j := 0, 0; j < len(s); {
		if !visited[s[j]] {
			visited[s[j]] = true
			j++
			l := j - i
			if l > ret {
				ret = l
			}
		} else {
			visited[s[i]] = false
			i++
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

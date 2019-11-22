package main

import "fmt"

func findAnagrams(s string, p string) []int {
	lp := len(p)
	if lp > len(s) {
		return []int{}
	}
	count := make([]int, 26)

	var i, j int
	for i = 0; i < lp; i++ {
		count[p[i]-'a']--
		count[s[i]-'a']++
	}

	ret := []int{}
	for i = 0; i < len(s)-lp+1; i++ {
		fmt.Println(count)
		for j = 0; j < 26; j++ {
			if count[j] != 0 {
				break
			}
		}
		if j == 26 {
			ret = append(ret, i)
		}
		if i+lp < len(s) {
			count[s[i]-'a']--
			count[s[i+lp]-'a']++
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

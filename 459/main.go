package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	count := make([]int, 26)
	for i := 0; i < n; i++ {
		count[s[i]-'a']++
	}

	minRepeat := n
	for i := 0; i < 26; i++ {
		if count[i] > 0 && count[i] < minRepeat {
			minRepeat = count[i]
		}
	}
	//fmt.Println(count)
	for k := 1; k < minRepeat; k++ {
		var i int
		if minRepeat%k != 0 {
			continue
		}
		countSub := minRepeat / k
		if n%countSub != 0 {
			continue
		}
		for i = 0; i < 26; i++ {
			if count[i]%countSub != 0 {
				break
			}
		}
		if i < 26 {
			continue
		}
		subLen := n / countSub
		for i = subLen; i < n; i++ {
			if s[i] != s[i%subLen] {
				break
			}
		}
		if i == n {
			//fmt.Println(k, subLen)
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("abaababaab"))
	fmt.Println(repeatedSubstringPattern("aabaaba"))
	fmt.Println(repeatedSubstringPattern("ababababababaababababababaababababababa"))
}

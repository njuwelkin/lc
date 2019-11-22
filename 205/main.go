package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var dicts [256]byte
	var dictt [256]byte
	for i := 0; i < len(s); i++ {
		if dicts[s[i]] == 0 && dictt[t[i]] == 0 {
			dicts[s[i]] = t[i]
			dictt[t[i]] = s[i]
		} else {
			if dicts[s[i]] != t[i] || dictt[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("ab", "aa"))
}

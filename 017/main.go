package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dict := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	buf := make([]byte, len(digits))
	ret := []string{}
	var f func(int)
	f = func(i int) {
		if i == len(digits) {
			ret = append(ret, string(buf))
			return
		}
		s := dict[digits[i]-'2']
		for j := 0; j < len(s); j++ {
			buf[i] = s[j]
			f(i + 1)
		}
	}
	f(0)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
}

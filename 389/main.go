package main

import "fmt"

func findTheDifference(s string, t string) byte {
	count := make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
		if count[t[i]-'a'] == -1 {
			return t[i]
		}
	}
	return 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%c\n", findTheDifference("abcd", "dcbae"))
}

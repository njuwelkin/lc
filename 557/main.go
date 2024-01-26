package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	buf := make([]byte, 0, len(s)+1)
	words := strings.Split(s, " ")
	for _, w := range words {
		for i := len(w) - 1; i >= 0; i-- {
			buf = append(buf, w[i])
		}
		buf = append(buf, ' ')
	}
	return string(buf[:len(buf)-1])
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

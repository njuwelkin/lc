package main

import (
	"fmt"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	buf := make([]byte, len(s))
	var i, j int
	for i = 0; i < len(s) && s[i] == ' '; i++ {
	}
	for j = 0; i < len(s); i++ {
		if i == 0 || s[i] != ' ' || s[i-1] != ' ' {
			buf[j] = s[i]
			j++
		}
	}
	if buf[j-1] == ' ' {
		j--
	}
	buf = buf[:j]

	reverse(buf)
	for i = 0; i < len(buf); {
		for j = i; j < len(buf) && buf[j] != ' '; j++ {
		}
		reverse(buf[i:j])
		for i = j; i < len(buf) && buf[i] == ' '; i++ {
		}
	}
	return string(buf)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("  hello    world!  "))
}

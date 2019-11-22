package main

import "fmt"

func compress(chars []byte) int {
	write := 0
	for i, j := 0, 1; j <= len(chars); j++ {
		if j == len(chars) || chars[j] != chars[i] {
			chars[write] = chars[i]
			write++
			l := j - i
			if l > 1 {
				str := fmt.Sprintf("%d", l)
				for k := 0; k < len(str); k++ {
					chars[write] = str[k]
					write++
				}
			}
			if j < len(chars) {
				i = j
			}
		}
	}
	fmt.Println(write, string(chars[:write]))
	return write
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(compress([]byte("aabbbcdd")))
}

package main

import "fmt"

func convertToBase7(num int) string {
	neg := false
	if num == 0 {
		return "0"
	} else if num < 0 {
		num = -num
		neg = true
	}
	buf := make([]byte, 128)
	idx := len(buf)
	for num > 0 {
		mod := num % 7
		num /= 7
		idx--
		buf[idx] = '0' + byte(mod)
	}
	if neg {
		idx--
		buf[idx] = '-'
	}
	return string(buf[idx:])
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(0))
}

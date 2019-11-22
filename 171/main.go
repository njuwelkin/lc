package main

import "fmt"

func titleToNumber(s string) int {
	ret := 0
	for _, c := range s {
		ret = 26*ret + int(c-'A') + 1
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

package main

import "fmt"

func countDigitOne(n int) int {
	ret := 0
	for n != 0 {
		n &= n - 1
		//n &= n - 1
		ret++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(countDigitOne(13))
	fmt.Println(countDigitOne(1))
	fmt.Println(countDigitOne(0))
}

package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}

	ret := 1
	for m != 0 {
		m >>= 1
		n >>= 1
		ret <<= 1
	}
	if n == 0 {
		return ret >> 1
	}

	return 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 1))
}

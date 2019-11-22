package main

import "fmt"

func divide(dividend int, divisor int) int {
	var d int
	for d = divisor; d < dividend; d <<= 1 {
	}
	d >>= 1

	ret := 0
	for ; d >= divisor; d >>= 1 {
		ret <<= 1
		if dividend > d {
			ret |= 1
			dividend -= d
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(divide(123, 9))
}

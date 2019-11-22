package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var i, j int
	for i, j = 0, x; i < j; {
		mid := (i + j) / 2
		if mid*mid <= x {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i - 1
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(0, mySqrt(0))
	fmt.Println(1, mySqrt(1))
	fmt.Println(2, mySqrt(2))
	fmt.Println(3, mySqrt(3))
	fmt.Println(4, mySqrt(4))
	fmt.Println(8, mySqrt(8))
	fmt.Println(9, mySqrt(9))
	fmt.Println(16, mySqrt(16))
}

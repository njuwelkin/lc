package main

import "fmt"

func hammingDistance(x int, y int) int {
	v := x ^ y
	ret := 0
	for v != 0 {
		ret++
		v &= v - 1
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(hammingDistance(1, 4))
}

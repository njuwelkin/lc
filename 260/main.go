package main

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	mask := 1
	for i := 0; i < 63; i++ {
		if xor&mask != 0 {
			break
		}
		mask <<= 1
	}
	c1, c2 := 0, 0
	for _, num := range nums {
		if num&mask != 0 {
			c1 ^= num
		} else {
			c2 ^= num
		}
	}
	return []int{c1, c2}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(singleNumber([]int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 7, 8, 8}))
}

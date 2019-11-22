package main

import "fmt"

func singleNumber(nums []int) int {
	countBit := make([]int, 64)
	for _, num := range nums {
		unum := uint(num)
		for i := 0; unum != 0; i++ {
			if unum&1 == 1 {
				countBit[i]++
			}
			unum >>= 1
		}
	}
	ret := 0
	mask := 1
	for i := 0; i < 64; i++ {
		if countBit[i]%3 == 1 {
			ret |= mask
		}
		mask <<= 1
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}

package main

import "fmt"

func countBit(num int, count []int) {
	for i := 0; num != 0; i++ {
		count[i] += num & 1
		num >>= 1
	}
}

func totalHammingDistance(nums []int) int {
	n := len(nums)
	count := make([]int, 64)
	for _, num := range nums {
		countBit(num, count)
	}

	ret := 0
	for i := 0; i < 64; i++ {
		ret += count[i] * (n - count[i])
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}

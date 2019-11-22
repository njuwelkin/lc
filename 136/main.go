package main

import "fmt"

func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
}

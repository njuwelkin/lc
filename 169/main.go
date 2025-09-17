package main

import "fmt"

func majorityElement(nums []int) int {
	val := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			val = num
		}
		if num == val {
			count++
		} else {
			count--
		}
	}
	return val
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

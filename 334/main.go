package main

import "fmt"

func increasingTriplet(nums []int) bool {
	first := 1 << 62
	second := 1 << 62
	for _, num := range nums {
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
}

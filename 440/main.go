package main

import "fmt"

func findKthNumber(n int, k int) int {
	count := 0
	ret := 0
	var helper func(num int) bool
	helper = func(num int) bool {
		if num > n {
			return false
		}
		count++
		fmt.Println(num, count)
		if count == k {
			ret = num
			return true
		}
		for i := 0; i < 10; i++ {
			if helper(num*10 + i) {
				return true
			}
		}
		return false
	}
	for i := 1; i < 10; i++ {
		if helper(i) {
			return ret
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findKthNumber(13, 2))
}

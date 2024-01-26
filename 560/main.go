package main

import "fmt"

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	sum := 0
	ret := 0
	for _, num := range nums {
		sum += num
		ret += m[sum-k]
		m[sum]++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1}, 0))
}

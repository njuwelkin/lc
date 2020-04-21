package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0, 1<<len(nums))
	ret = append(ret, []int{})
	for i := 0; i < len(nums); {
		var j int
		num := nums[i]
		for _, exist := range ret {
			newItems := []int{}
			for j = i; j < len(nums) && nums[j] == num; j++ {
				newItems = append(newItems, num)
				buf := make([]int, len(exist), len(exist)+len(newItems))
				copy(buf, exist)
				buf = append(buf, newItems...)
				ret = append(ret, buf)
			}
		}
		i = j
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(subsets([]int{1, 2, 3}))
	//fmt.Println(subsets([]int{1, 1, 2, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	//fmt.Println(subsets([]int{0, 3, 5, 7}))
}

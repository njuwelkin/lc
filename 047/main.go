package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	ret := [][]int{{}}
	for i := 0; i < n; {
		num := nums[i]
		next := [][]int{}

		var j int
		for j = i + 1; nums[j] != nums[i]; j++ {
		}
		count := j - i
		i = j

		for _, p := range ret {
			tmp := make([]int, len(p)+count)
			for k := len(p); k < len(tmp); k++ {
				tmp[k] = num
			}
			next = append(next, tmp)
		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

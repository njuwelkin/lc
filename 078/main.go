package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	ret := make([][]int, 1<<uint(len(nums)))
	retIdx := 0
	appendRet := func(val []int) {
		ret[retIdx] = val
		retIdx++
	}
	appendRet([]int{})

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j := retIdx - 1; j >= 0; j-- {
			//newItem := make([]int, len(ret[j]))
			//copy(newItem, ret[j])
			newItem := append(ret[j], nums[i])
			appendRet(newItem)
		}
	}

	//debug
	sort.Slice(ret, func(i, j int) bool {
		if len(ret[i]) == len(ret[j]) {
			for k := 0; k < len(ret[i]); k++ {
				if ret[i][k] < ret[j][k] {
					return true
				} else if ret[i][k] > ret[j][k] {
					return false
				}
			}
			return false
		}
		return len(ret[i]) < len(ret[j])
	})
	fmt.Println(retIdx)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

package main

import (
	"fmt"
	"sort"
)

func equal(n1, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			return false
		}
	}
	return true
}

func findSubsequences(nums []int) [][]int {
	ret := [][]int{{}}

	for _, num := range nums {
		retlen := len(ret)
		for j := 0; j < retlen; j++ {
			if len(ret[j]) == 0 || num >= ret[j][len(ret[j])-1] {
				tmp := make([]int, len(ret[j]))
				copy(tmp, ret[j])
				tmp = append(tmp, num)
				ret = append(ret, tmp)
			}
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		if len(ret[i]) == len(ret[j]) {
			for k := 0; k < len(ret[i]); k++ {
				if ret[i][k] < ret[j][k] {
					return true
				} else if ret[i][k] < ret[j][k] {
					return false
				}
			}
		}
		return len(ret[i]) < len(ret[j])
	})

	var i, j int
	for i, j = 0, 0; j < len(ret); j++ {
		if len(ret[j]) > 1 && !equal(ret[j], ret[j-1]) {
			ret[i] = ret[j]
			i++
		}
	}
	return ret[:i]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

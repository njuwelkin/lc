package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		if _, found := m[num]; !found {
			m[num] = 1
		} else {
			m[num]++
		}
	}

	bucket := make([][]int, len(nums)+1)
	for num, count := range m {
		bucket[count] = append(bucket[count], num)
	}
	ret := []int{}
	for i := len(nums); i > 0; i-- {
		for _, num := range bucket[i] {
			ret = append(ret, num)
			k--
			if k == 0 {
				return ret
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{}, 2))
}

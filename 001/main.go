package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int][]int{}
	for i, num := range nums {
		if _, found := m[num]; !found {
			m[num] = []int{1, i}
		} else {
			m[num] = []int{2, i}
		}
	}

	for i, num := range nums {
		if v, found := m[target-num]; found {
			if target-num != num || v[0] == 2 {
				return []int{i, v[1]}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

package main

import "fmt"

func findPairs(nums []int, k int) int {
	m := map[int]int{}
	for _, num := range nums {
		if count, found := m[num]; !found {
			m[num] = 1
		} else {
			m[num] = count + 1
		}
	}
	ret := 0
	if k != 0 {
		for key, _ := range m {
			if _, found := m[key-k]; found {
				ret += 1
			}
		}
	} else {
		for _, v := range m {
			if v > 1 {
				ret += 1
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}

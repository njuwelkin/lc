package main

import "fmt"

func longestConsecutive(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] = 1
	}

	ret := 0
	for _, num := range nums {
		if _, found := m[num-1]; !found {
			count := 1
			for {
				num++
				if _, found := m[num]; found {
					count++
					m[num] = count
				} else {
					break
				}
			}
			if count > ret {
				ret = count
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

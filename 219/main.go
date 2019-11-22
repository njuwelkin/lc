package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if j, found := m[num]; found {
			if i-j <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

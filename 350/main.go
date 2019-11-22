package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	ret := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num]--
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{2, 6, 2, 9, 1}, []int{7, 1}))
}

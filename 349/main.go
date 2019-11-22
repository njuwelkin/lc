package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, num := range nums1 {
		m[num] = true
	}

	ret := []int{}
	for _, num := range nums2 {
		if m[num] {
			ret = append(ret, num)
			m[num] = false
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	//fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersection([]int{2, 6, 2, 9, 1}, []int{7, 1}))
}

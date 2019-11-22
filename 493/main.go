package main

import "fmt"

func merge(n []int, mid int) int {
	n1 := make([]int, mid)
	copy(n1, n[:mid])
	n2 := n[mid:]

	ret := 0
	var i, j, k int
	for i, j = 0, 0; i < len(n1); {
		if j < len(n2) && n1[i] > 2*n2[j] {
			j++
		} else {
			i++
			ret += j
		}
	}

	// sort
	for i, j, k = 0, 0, 0; i < len(n1) && j < len(n2); {
		if n1[i] <= n2[j] {
			n[k] = n1[i]
			i++
		} else {
			n[k] = n2[j]
			j++
		}
		k++
	}
	for ; i < len(n1); i++ {
		n[k] = n1[i]
		k++
	}
	return ret
}

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	ret := reversePairs(nums[:mid]) + reversePairs(nums[mid:])
	ret += merge(nums, mid)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
}

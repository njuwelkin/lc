package main

import "fmt"

func max(nums []int) int {
	idx := -1
	tmp := -1
	for i, num := range nums {
		if num > tmp {
			idx = i
		}
	}
	return idx
}

func getCandidateRange(nums1, nums2 []int, k int) (int, int) {
	var c1, c2 int
	if len(nums2) >= k-1 {
		c1 = len(nums1)
	} else {
		c1 = len(num1) - (k - 1 - len(nums2))
	}
	if len(nums1) >= k-1 {
		c2 = len(nums2)
	} else {
		c2 = len(num2) - (k - 1 - len(nums1))
	}
	return c1, c2
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ret := make([]int, 0, k)

	c1, c2 := getCandidateRange(nums1, nums2, 0, 0, k)
	i, j := start1+max(nums1[start1:c1]), start2+max(nums2[start2:c2])
	for start1, start2 := 0, 0; start1 < len(nums1) && start2 < len(nums2); {
		k--
		if start2 >= len(nums2) || nums1[i] > nums2[j] {
			ret = append(ret, nums1[i])
			start1 = i + 1
			c1, c2 = getCandidateRange(nums1[start1:], nums2[start2:], k
			c1 += start1
			c2 += start2
		} else if nums1[i] < nums2[j] {
			ret = append(ret, nums2[j])
			start2 = j + 1
			c1, c2 = getCandidateRange(nums1[start1:], nums2[start2:], k)
			c1 += start1
			c2 += start2
		} else {

		}
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{8, 1, 2, 5, 9, 3}, 5))
}

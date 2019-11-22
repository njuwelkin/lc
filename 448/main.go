package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	tag := 1 << 62
	mask := tag - 1
	for _, num := range nums {
		nums[mask&num-1] |= tag
	}
	ret := []int{}
	for i, num := range nums {
		if num&tag == 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

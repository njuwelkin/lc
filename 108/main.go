package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(sortedArrayToBST([]int{}).ToString())
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}).ToString())
}

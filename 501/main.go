package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func iterate(root *TreeNode, f func(*TreeNode)) {
	var midOrder func(root *TreeNode)
	midOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		midOrder(root.Left)
		f(root)
		midOrder(root.Right)
	}
	midOrder(root)
}

func findMode(root *TreeNode) []int {
	maxCount := 0
	count := [][]int{}
	prev := -1 << 63
	iterate(root, func(p *TreeNode) {
		//fmt.Println(p.Val)
		if p.Val != prev {
			count = append(count, []int{p.Val, 1})
			prev = p.Val
		} else {
			count[len(count)-1][1]++
		}
		if count[len(count)-1][1] > maxCount {
			maxCount = count[len(count)-1][1]
		}
	})

	fmt.Println(count, maxCount)
	ret := []int{}
	for j := 0; j < len(count); j++ {
		if count[j][1] == maxCount {
			ret = append(ret, count[j][0])
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(findMode(NewBTree([]int{3, 1, 6, 1, 2, 5, 7, Null, Null, Null, Null, 5})))
	fmt.Println(findMode(NewBTree([]int{1, Null, 2, 2})))
}

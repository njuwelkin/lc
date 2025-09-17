package main

import (
	"fmt"

	. "github.com/njuwelkin/lc/ds"
)

func depth(root *TreeNode, path int, height uint) int {
	ret := 0
	mask := 1 << (height - 2)
	for p := root; p != nil; {
		ret++
		if path&mask == 0 {
			p = p.Left
		} else {
			p = p.Right
		}
		mask >>= 1
	}
	return ret
}

func countNodes(root *TreeNode) int {
	maxDepth, minDepth := 0, 0
	for p := root; p != nil; p = p.Left {
		maxDepth++
	}
	for p := root; p != nil; p = p.Right {
		minDepth++
	}
	if maxDepth == minDepth {
		return 1<<uint(maxDepth) - 1
	}

	var i, j int
	for i, j = 0, 1<<uint(maxDepth-1); i < j; {
		mid := (i + j) / 2
		d := depth(root, mid, uint(maxDepth))
		if d == maxDepth {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return 1<<uint(minDepth) - 1 + i
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(countNodes(NewBTree([]int{1, 2, 3, 4, 5, 6, Null})))
	//fmt.Println(countNodes(nil))
	//fmt.Println(countNodes(NewBTree([]int{1, Null, Null})))
	fmt.Println(countNodes(NewBTree([]int{1, 2, 3, 4})))
}

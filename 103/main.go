package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	crtLevel := []*TreeNode{root}

	level := 0
	for len(crtLevel) != 0 {
		line := make([]int, len(crtLevel))
		nextLevel := []*TreeNode{}

		for i := len(crtLevel) - 1; i >= 0; i-- {
			p := crtLevel[i]
			line[len(line)-i-1] = p.Val
			if level%2 == 0 {
				if p.Left != nil {
					nextLevel = append(nextLevel, p.Left)
				}
				if p.Right != nil {
					nextLevel = append(nextLevel, p.Right)
				}
			} else {
				if p.Right != nil {
					nextLevel = append(nextLevel, p.Right)
				}
				if p.Left != nil {
					nextLevel = append(nextLevel, p.Left)
				}
			}
		}
		ret = append(ret, line)
		crtLevel = nextLevel
		level++
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(zigzagLevelOrder(NewBTree([]int{3, 9, 20, Null, Null, 15, 7})))
}

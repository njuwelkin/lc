package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func sortList(head *ListNode) *ListNode {
	dummyHead := ListNode{1 << 62, head}
	for p := &dummyHead; p.Next != nil; {
		prev := &dummyHead
		min := 1 << 62
		for q := p; q.Next != nil; q = q.Next {
			if q.Next.Val < min {
				min = q.Next.Val
				prev = q
			}
		}
		if prev != p {
			minNode := prev.Next
			prev.Next = minNode.Next
			minNode.Next = p.Next
			p.Next = minNode
		}
		p = p.Next
	}
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(sortList(NewList([]int{-1, 5, 3, 4, 0})).ToArray())
}

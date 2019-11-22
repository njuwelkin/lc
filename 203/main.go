package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := ListNode{0, head}
	for p := &dummyHead; p != nil && p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		}
	}
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(removeElements(NewList([]int{1, 2, 6, 3, 4, 5, 6}), 6).ToArray())
	fmt.Println(removeElements(NewList([]int{1, 2, 6, 3, 4, 5, 6}), 1).ToArray())
}

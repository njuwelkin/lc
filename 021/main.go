package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	p := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return head.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(mergeTwoLists(NewList([]int{1, 3, 5, 7, 9}), NewList([]int{2, 4, 6})).ToArray())
	fmt.Println(mergeTwoLists(NewList([]int{1, 3, 5, 7, 9}), nil).ToArray())
}

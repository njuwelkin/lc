package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func merge(l1, l2 *ListNode) *ListNode {
	head := ListNode{}
	p := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	return merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(mergeKLists([]*ListNode{NewList([]int{1, 4, 5}), NewList([]int{1, 3, 4}), NewList([]int{2, 6})}).ToArray())
}

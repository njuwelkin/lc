package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := ListNode{Next: head}
	p := &dummyHead
	for i := 0; i < n; i++ {
		if p == nil {
			break
		}
		p = p.Next
	}
	if p == nil {
		return head
	}
	p = p.Next
	q := &dummyHead
	for p != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(removeNthFromEnd(NewList([]int{1, 2, 3, 4, 5}), 2).ToArray())
	fmt.Println(removeNthFromEnd(NewList([]int{1, 2, 3, 4, 5}), 6).ToArray())
}

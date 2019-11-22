package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func merge(p, q *ListNode) *ListNode {
	dummyHead := ListNode{}
	tail := &dummyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			tail = p
			p = p.Next
		} else {
			tail.Next = q
			tail = q
			q = q.Next
		}
	}
	if p != nil {
		tail.Next = p
	} else {
		tail.Next = q
	}
	return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	p := head
	for i := 0; i < count/2-1; i++ {
		p = p.Next
	}
	q := p.Next
	p.Next = nil
	return merge(sortList(head), sortList(q))
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(sortList(NewList([]int{-1, 5, 3, 4, 0})).ToArray())
}

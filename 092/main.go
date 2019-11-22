package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead := ListNode{0, head}
	p := &dummyHead
	var i int
	for i = 0; i < m-1 && p != nil; i++ {
		p = p.Next
	}
	if p == nil {
		return head
	}
	revHead := p
	q := p.Next

	for ; i < n && p != nil; i++ {
		p = p.Next
	}
	if p == nil {
		revHead.Next = nil
	} else {
		fmt.Println(p.Val)
		revHead.Next = p.Next
		p.Next = nil
	}

	for q != nil {
		next := q.Next
		q.Next = revHead.Next
		revHead.Next = q
		q = next
	}
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reverseBetween(NewList([]int{1, 2, 3, 4, 5}), 2, 4).ToArray())
	fmt.Println(reverseBetween(NewList([]int{1, 2, 3, 4, 5}), 1, 4).ToArray())
	fmt.Println(reverseBetween(NewList([]int{1, 2, 3, 4, 5}), 3, 8).ToArray())
	fmt.Println(reverseBetween(NewList([]int{}), 2, 4).ToArray())
}

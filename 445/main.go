package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func reverse(l *ListNode) (*ListNode, int) {
	head := ListNode{}
	count := 0
	for l != nil {
		tmp := l.Next
		l.Next = head.Next
		head.Next = l
		l = tmp
		count++
	}
	return head.Next, count
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, n1 := reverse(l1)
	l2, n2 := reverse(l2)
	if n2 > n1 {
		l1, l2 = l2, l1
	}
	if n1 == 0 {
		return nil
	}

	prev := 0
	var p, q, tail *ListNode
	for p, q = l1, l2; p != nil; p = p.Next {
		p.Val += prev
		if q != nil {
			p.Val += q.Val
			q = q.Next
		}
		prev = p.Val / 10
		p.Val %= 10
		tail = p
	}
	if prev > 0 {
		tail.Next = &ListNode{Val: 1}
	}
	ret, _ := reverse(l1)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(addTwoNumbers(NewList([]int{7, 2, 4, 3}), NewList([]int{5, 6, 4})).ToArray())
}

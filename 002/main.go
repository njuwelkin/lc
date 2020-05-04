package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Next: l1}

	tail := head
	carry := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		if val >= 10 {
			carry = 1
			val %= 10
		} else {
			carry = 0
		}
		l1.Val = val

		tail = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		tail.Next = l2
		l1 = l2
	}

	for l1 != nil && carry == 1 {
		l1.Val += carry
		if l1.Val >= 10 {
			l1.Val %= 10
		} else {
			carry = 0
		}
		tail = l1
		l1 = l1.Next
	}

	if carry == 1 {
		tail.Next = &ListNode{Val: 1}
	}
	return head.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(addTwoNumbers(NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4})).ToArray())
	fmt.Println(addTwoNumbers(NewList([]int{2, 4, 3}), NewList([]int{5, 6, 6, 9})).ToArray())
}

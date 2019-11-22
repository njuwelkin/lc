package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func reverse(head *ListNode) *ListNode {
	dummy := ListNode{0, nil}
	for p := head; p != nil; {
		tmp := p.Next
		p.Next = dummy.Next
		dummy.Next = p
		p = tmp
	}
	return dummy.Next
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var p, q *ListNode
	for p, q = head, head.Next; ; {
		q = q.Next
		if q == nil {
			break
		}
		q = q.Next
		if q == nil {
			break
		}
		p = p.Next
	}
	q = reverse(p.Next)
	p.Next = nil
	fmt.Println(head.ToArray(), q.ToArray())
	for p = head; p != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isPalindrome(NewList([]int{1, 2, 3, 2, 1})))
	fmt.Println(isPalindrome(NewList([]int{1, 2, 2, 1})))
}

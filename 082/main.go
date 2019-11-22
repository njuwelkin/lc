package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := ListNode{0, head}
	for p := &fakeHead; p != nil; {
		if p.Next == nil {
			break
		}
		val := p.Next.Val
		var q *ListNode
		for q = p.Next; q != nil && q.Val == val; q = q.Next {
		}
		if q == p.Next.Next {
			p = p.Next
		} else {
			p.Next = q
		}
	}
	return fakeHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(deleteDuplicates(NewList([]int{1, 2, 3, 3, 4, 4, 5})).ToArray())
	fmt.Println(deleteDuplicates(NewList([]int{1, 1, 1, 2, 3})).ToArray())
}

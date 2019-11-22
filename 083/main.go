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
	for p := head; p.Next != nil; {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(deleteDuplicates(NewList([]int{1, 2, 3, 3, 4, 4, 5})).ToArray())
	fmt.Println(deleteDuplicates(NewList([]int{1, 1, 1, 2, 3})).ToArray())
}

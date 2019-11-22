package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	count := 0
	p := head
	for ; p.Next != nil; p = p.Next {
		count++
	}
	count++
	k = (count - k%count) % count
	p.Next = head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	ret := p.Next
	p.Next = nil
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(rotateRight(NewList([]int{1, 2, 3, 4, 5}), 2).ToArray())
	fmt.Println(rotateRight(NewList([]int{1, 2, 3, 4, 5}), 8).ToArray())
}

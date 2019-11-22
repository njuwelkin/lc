package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := ListNode{0, head}
	for p := &dummyHead; p != nil; {
		first := p.Next
		if first == nil {
			break
		}
		second := first.Next
		if second == nil {
			break
		}
		p.Next = second
		first.Next = second.Next
		second.Next = first
		p = first
	}
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(swapPairs(NewList([]int{1, 2, 3, 4})).ToArray())
	fmt.Println(swapPairs(NewList([]int{1, 2, 3, 4, 5})).ToArray())
}

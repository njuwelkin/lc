package main

import (
	"fmt"
	. "github.com/njuwelkin/lc/ds"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := ListNode{0, head}
	for p := &dummyHead; ; {
		nextP := p.Next
		q := p
		for i := 0; i < k && q != nil; i++ {
			q = q.Next
		}
		if q == nil {
			break
		}
		nextQ := q.Next
		q.Next = nil
		for q = nextP; q != nil; {
			next := q.Next
			q.Next = p.Next
			p.Next = q
			q = next
		}
		p = nextP
		p.Next = nextQ
	}
	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reverseKGroup(NewList([]int{1, 2, 3, 4, 5}), 3).ToArray())
}

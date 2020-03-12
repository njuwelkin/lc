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
type Solution struct {
	head   *ListNode
	length int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	return Solution{head, count}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor(NewList([]int{1, 2, 3, 4, 5}))
	fmt.Println(obj.GetRandom())
}

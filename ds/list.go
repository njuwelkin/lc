package ds

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, num := range nums {
		p.Next = &ListNode{num, nil}
		p = p.Next
	}
	return head.Next
}

func (l *ListNode) ToArray() []int {
	ret := []int{}
	for p := l; p != nil; p = p.Next {
		ret = append(ret, p.Val)
	}
	return ret
}

package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

const Null = -1

func NewList(items [][2]int) *Node {
	i2p := map[int]*Node{}
	head := Node{}
	p := &head
	for i, item := range items {
		p.Next = &Node{item[0], nil, nil}
		p = p.Next
		i2p[i] = p
	}
	fmt.Println(i2p)
	for i, p := 0, head.Next; p != nil; p = p.Next {
		p.Random = i2p[items[i][1]]
		i++
		if p.Random != nil {
			fmt.Println(p.Val, p.Random.Val)
		} else {
			fmt.Println(p.Val)
		}
	}
	return head.Next
}

func (l *Node) ToArray() [][2]int {
	p2i := map[*Node]int{}

	ret := [][2]int{}
	for p := l; p != nil; p = p.Next {
		ret = append(ret, [2]int{p.Val, -1})
		p2i[p] = len(ret) - 1
	}
	fmt.Println(p2i)
	i := 0
	for p := l; p != nil; p = p.Next {
		if p.Random != nil {
			ret[i][1] = p2i[p.Random]
		}
		i++
	}
	return ret
}

func copyRandomList(head *Node) *Node {
	dummyHead := Node{}
	q := &dummyHead
	for p := head; p != nil; {
		next := p.Next
		q.Next = &Node{
			Val:    p.Val,
			Random: p,
		}
		q = q.Next
		p.Next = q
		p = next
	}
	fmt.Println(dummyHead.Next.ToArray())

	for p := dummyHead.Next; p != nil; p = p.Next {
		if p.Random.Random != nil {
			p.Random = p.Random.Random.Next
		}
	}

	return dummyHead.Next
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(NewList([][2]int{{7, Null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}).ToArray())
	fmt.Println(copyRandomList(NewList([][2]int{{7, Null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})).ToArray())
}

package ds

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const Null = math.MaxInt64

func NewBTree(vals []int) *TreeNode {
	root, _ := DeserializeBTree(vals)
	return root
}

func DeserializeBTree(vals []int) (*TreeNode, error) {
	if len(vals) == 0 {
		return nil, nil
	}
	root := TreeNode{vals[0], nil, nil}
	q := NewQueue()
	q.Add(&root)
	i := 1
	for !q.IsEmpty() && i < len(vals) {
		if i >= len(vals) {
			break
		}
		node, _ := q.Delete()
		if vals[i] != Null {
			node.(*TreeNode).Left = &TreeNode{vals[i], nil, nil}
			q.Add(node.(*TreeNode).Left)
		}
		i++
		if i < len(vals) && vals[i] != Null {
			node.(*TreeNode).Right = &TreeNode{vals[i], nil, nil}
			q.Add(node.(*TreeNode).Right)
		}
		i++
	}
	return &root, nil
}

func (root *TreeNode) Serialize() []int {
	if root == nil {
		return []int{}
	}
	ret := []int{root.Val}
	q := NewQueue()
	q.Add(root)
	for !q.IsEmpty() {
		node, _ := q.Delete()
		btNode := node.(*TreeNode)
		left := btNode.Left
		right := btNode.Right
		if left != nil {
			q.Add(left)
			ret = append(ret, left.Val)
		} else {
			ret = append(ret, Null)
		}
		if right != nil {
			q.Add(right)
			ret = append(ret, right.Val)
		} else {
			ret = append(ret, Null)
		}
	}
	for i := len(ret) - 1; i >= 0; i-- {
		if ret[i] != Null {
			return ret[:i+1]
		}
	}
	return ret
}

func (root *TreeNode) ToString() string {
	ret := "["
	array := root.Serialize()
	for i, val := range array {
		if i != 0 {
			ret += ", "
		}
		if val == Null {
			ret += "null"
		} else {
			ret += fmt.Sprintf("%d", val)
		}
	}
	ret += "]"
	return ret
}

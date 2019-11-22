package main

import "fmt"

type NumArray struct {
	tree []int
}

func left(i int) int {
	return 2 * i
}
func right(i int) int {
	return 2*i + 1
}
func parent(i int) int {
	return i / 2
}

func Constructor(nums []int) NumArray {
	treeSize := 1
	for treeSize < len(nums) {
		treeSize <<= 1
	}
	treeSize <<= 1
	ret := NumArray{make([]int, treeSize)}
	for i := 0; i < len(nums); i++ {
		ret.tree[i+treeSize/2] = nums[i]
	}
	for i := treeSize - 1; i > 1; i-- {
		ret.tree[parent(i)] += ret.tree[i]
	}
	return ret
}

func (this *NumArray) getIdxInTree(idx int) int {
	return len(this.tree)/2 + idx
}

func (this *NumArray) Update(i int, val int) {
	idx := this.getIdxInTree(i)
	d := val - this.tree[idx]
	for ; idx != 0; idx = parent(idx) {
		this.tree[idx] += d
	}
}

func (this *NumArray) sumLeft(i int) int {
	ret := 0
	for idx := this.getIdxInTree(i); idx != 1; idx = parent(idx) {
		p := parent(idx)
		if idx == right(p) {
			ret += this.tree[left(p)]
		}
	}
	return ret
}

func (this *NumArray) sumRight(i int) int {
	ret := 0
	for idx := this.getIdxInTree(i); idx != 1; idx = parent(idx) {
		p := parent(idx)
		if idx == left(p) {
			ret += this.tree[right(p)]
		}
	}
	return ret
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.tree[1] - this.sumLeft(i) - this.sumRight(j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor([]int{1, 3, 5, 7, 9})
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 3)
	fmt.Println(obj.SumRange(1, 3))
	obj.Update(2, 3)
	fmt.Println(obj.SumRange(1, 3))
}

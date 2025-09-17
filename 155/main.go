package main

import "fmt"

type MinStack struct {
	min []int
	val []int
}

func Constructor() MinStack {
	return MinStack{
		min: []int{1 << 62},
		val: []int{1 << 62},
	}
}

func (this *MinStack) Push(val int) {
	this.val = append(this.val, val)
	if min := this.min[len(this.min)-1]; val < min {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.val) > 1 {
		this.val = this.val[0 : len(this.val)-1]
		this.min = this.val[0 : len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	fmt.Println("vim-go")
}

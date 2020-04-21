package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	ret := Solution{map[int][]int{}}
	for i, num := range nums {
		ret.m[num] = append(ret.m[num], i)
	}
	return ret
}

func (this *Solution) Pick(target int) int {
	idx := this.m[target]
	if idx == nil || len(idx) == 0 {
		return -1
	}
	return idx[rand.Intn(len(idx))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(obj.m[123] == nil)
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(3))
	fmt.Println(obj.Pick(1))
	fmt.Println(obj.Pick(0))
}

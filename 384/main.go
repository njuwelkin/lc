package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	ret := make([]int, n)
	copy(ret, this.nums)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(obj.Reset())
	fmt.Println(obj.Shuffle())
	fmt.Println(obj.Reset())
	fmt.Println(obj.Shuffle())
}

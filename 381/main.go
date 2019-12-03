package main

import (
	"fmt"
	"math/rand"
)

type RandomizedCollection struct {
	m map[int][]int
	a []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m: make(map[int][]int),
		a: []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	var found bool
	if _, found = this.m[val]; !found {
		this.m[val] = []int{}
	}
	this.m[val] = append(this.m[val], len(this.a))
	this.a = append(this.a, val)
	return !found
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	var found bool
	if _, found = this.m[val]; !found {
		return false
	}
	if len(this.a) == 1 {
		this.a = []int{}
		this.m = map[int]int{}
	} else {
		idx := this.m[val]
		val2 := this.a[len(this.a)-1]
		this.a[idx] = val2
		this.a = this.a[:len(this.a)-1]
		this.m[val2] = idx
		delete(this.m, val)
	}
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	fmt.Println("vim-go")
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	obj.Remove(3)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

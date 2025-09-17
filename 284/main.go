package main

import "fmt"

/*   Below is the interface for Iterator, which is already defined for you.
 */
type Iterator struct {
	nums []int
	idx  int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return this.idx < len(this.nums)
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	ret := this.nums[this.idx]
	this.idx++
	return ret
}

type PeekingIterator struct {
	iter       *Iterator
	current    int
	terminated bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	ret := PeekingIterator{
		iter:       iter,
		terminated: !iter.hasNext(),
	}
	if !ret.terminated {
		ret.current = iter.next()
	}
	return &ret
}

func (this *PeekingIterator) hasNext() bool {
	return !this.terminated
}

func (this *PeekingIterator) next() int {
	ret := this.current
	if this.iter.hasNext() {
		this.current = this.iter.next()
	} else {
		this.terminated = true
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this.current
}

func main() {
	fmt.Println("vim-go")
}

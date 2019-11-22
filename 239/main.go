package main

import "fmt"

type Win struct {
	buf        []int
	head, tail int
}

func NewWin(n int) Win {
	return Win{make([]int, n), 0, 0}
}

func (w *Win) Insert(num int) {
	i, j := w.head, w.tail
	if j <= i {
		j += len(w.buf)
	}
	for i < j {
		mid := (i + j) / 2
		if w.buf[mid%len(w.buf)] > num {
			i = mid + 1
		} else {
			j = mid
		}
	}
	w.buf[i%len(w.buf)] = num
	w.tail = (i + 1) % len(w.buf)
}

func (w *Win) Delete(num int) {
	if w.buf[w.head] == num {
		w.head = (w.head + 1) % len(w.buf)
	}
}

func (w *Win) Top() int {
	return w.buf[w.head]
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 || len(nums) < k {
		return []int{}
	}
	w := NewWin(k)
	for i := 0; i < k; i++ {
		w.Insert(nums[i])
	}
	ret := []int{w.Top()}

	for i := k; i < len(nums); i++ {
		w.Delete(nums[i-k])
		w.Insert(nums[i])
		ret = append(ret, w.Top())
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 1))
}

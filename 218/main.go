package main

import (
	"fmt"
	"math"
	"sort"
)

type Stack [][]int

func (s *Stack) Push(x []int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() []int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s Stack) Top() []int {
	return s[len(s)-1]
}

func (s Stack) Ajust() {
	for i := len(s) - 1; i > 0; i-- {
		if s[i][2] > s[i-1][2] {
			break
		}
		s[i], s[i-1] = s[i-1], s[i]
	}
}

func getSkyline(buildings [][]int) [][]int {
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][0] == buildings[j][0] {
			if buildings[i][2] == buildings[j][2] {
				return buildings[i][1] > buildings[j][1]
			}
			return buildings[i][2] > buildings[j][2]
		}
		return buildings[i][0] < buildings[j][0]
	})

	ret := [][]int{}
	appendRet := func(x, y int) {
		ret = append(ret, []int{x, y})
	}

	stack := Stack([][]int{})
	stack.Push([]int{math.MinInt64, math.MaxInt64, 0})
	for i := 0; i < len(buildings); {
		fmt.Println(stack)
		b := buildings[i]
		top := stack.Top()
		if b[0] == top[0] { //top.left == b.left
			if b[2] == top[2] { // height
				i++
				continue
			} else {
				if b[1] <= top[1] {
					i++
					continue
				} else {
					stack.Push(b)
					stack.Ajust()
					i++
				}
			}
		} else { // top.left < b.left
			if b[0] <= top[1] { // overlap
				if b[1] <= top[1] { // include
					if b[2] <= top[2] { // prev building is taller
						i++
						continue
					} else { // new building is taller
						appendRet(b[0], b[2])
						stack.Push(b)
						i++
					}
				} else {
					if b[2] == top[2] {
						top[1] = b[1]
						i++
					} else if b[2] > top[2] {
						appendRet(b[0], b[2])
						stack.Push(b)
						i++
					} else {
						stack.Push(b)
						stack.Ajust()
						i++
					}
				}
			} else { // no overlap
				stack.Pop()
				if ret[len(ret)-1][0] < top[1] {
					var j int
					for j = len(stack) - 1; j >= 0 && stack[j][1] <= top[1]; j-- {
					}
					appendRet(top[1], stack[j][2])
				}
			}
		}
	}
	for len(stack) > 1 {
		fmt.Println(stack)
		top := stack.Pop()
		if ret[len(ret)-1][0] < top[1] {
			var j int
			for j = len(stack) - 1; j >= 0 && stack[j][1] <= top[1]; j-- {
			}
			appendRet(top[1], stack[j][2])
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
	fmt.Println(getSkyline([][]int{{3, 7, 8}, {3, 8, 7}, {3, 9, 6}, {3, 10, 5}, {3, 11, 4}, {3, 12, 3}, {3, 13, 2}, {3, 14, 1}}))
}

package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	rects [][]int
	area  []int
}

func area(rect []int) int {
	return (rect[2] + 1 - rect[0]) * (rect[3] + 1 - rect[1])
}

func Constructor(rects [][]int) Solution {
	ret := Solution{make([][]int, len(rects)),
		make([]int, len(rects))}
	ret.rects[0] = []int{rects[0][0], rects[0][1],
		rects[0][2] + 1 - rects[0][0]} // left, top, width
	ret.area[0] = area(rects[0])
	for i := 1; i < len(rects); i++ {
		rect := rects[i]
		ret.rects[i] = []int{rect[0], rect[1], rect[2] + 1 - rect[0]}
		ret.area[i] = ret.area[i-1] + area(rect)
	}

	return ret
}

func (this *Solution) findRect(area int) int {
	i, j := 0, len(this.area)
	for i < j {
		mid := (i + j) / 2
		if area >= this.area[mid] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func (this *Solution) findCoordinate(i int, area int) []int {
	rect := this.rects[i]
	if i != 0 {
		area = area - this.area[i-1]
	}
	ret := make([]int, 2)
	width := rect[2]
	ret[0] = rect[0] + area%width
	ret[1] = rect[1] + area/width
	return ret
}

func (this *Solution) Pick() []int {
	x := rand.Intn(this.area[len(this.area)-1])
	i := this.findRect(x)
	fmt.Println(x, i)
	return this.findCoordinate(i, x)
}

func verify(rects [][]int, obj Solution, n int) {
	fmt.Println(obj)
	for i := 0; i < n; i++ {
		point := obj.Pick()
		found := false
		for _, rect := range rects {
			if point[0] >= rect[0] && point[0] <= rect[2] &&
				point[1] >= rect[1] && point[1] <= rect[3] {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("error: point %v invalid\n", point)
			break
		}
	}
}

func main() {
	fmt.Println("vim-go")
	//rects := [][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}}
	rects := [][]int{{82918473, -57180867, 82918476, -57180863}, {83793579, 18088559, 83793580, 18088560}, {66574245, 26243152, 66574246, 26243153}, {72983930, 11921716, 72983934, 11921720}}
	obj := Constructor(rects)
	verify(rects, obj, 10000)
}

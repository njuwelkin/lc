package main

import (
	"fmt"
	"math"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	heaters = append(heaters, math.MaxInt64)

	ret := 0
	for i, j := 0, 0; i < len(houses); i++ {
		for ; heaters[j] < houses[i]; j++ {
		}
		if j == 0 {
			ret = max(ret, heaters[j]-houses[i])
		} else {
			ret = max(ret, min(heaters[j]-houses[i], houses[i]-heaters[j-1]))
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
}

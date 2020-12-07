package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	dp := map[[2]int][][]int{}

	var f func(c []int, t int) [][]int
	f = func(c []int, t int) [][]int {
		//fmt.Println(c, t)
		if r, found := dp[[2]int{t, len(c)}]; found {
			return r
		}
		if t == 0 {
			return [][]int{[]int{}}
		}
		if len(c) == 0 {
			return [][]int{}
		}
		ret := [][]int{}
		for i, num := range c {
			if num > t {
				break
			}

			//fmt.Println(num)
			res2 := f(c[i:], t-num)
			for _, item := range res2 {
				ret = append(ret, append(item, num))
			}
		}
		dp[[2]int{t, len(c)}] = ret
		//fmt.Println(dp)
		return ret
	}
	return f(candidates, target)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equal2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, itemA := range a {
		found := false
		for _, itemB := range b {
			if equal(itemA, itemB) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func combinationSum(candidates []int, target int) [][]int {
	dp := map[[2]int][][]int{}
	sort.Ints(candidates)

	var f func([]int, int) [][]int
	f = func(cand []int, tar int) [][]int {
		if val, visited := dp[[2]int{len(cand), tar}]; visited {
			fmt.Println(len(cand), tar)
			fmt.Println(dp[[2]int{len(cand), tar}])
			return val
		}
		if tar == 0 {
			return [][]int{[]int{}}
		}
		var i int
		for i = len(cand) - 1; i >= 0 && cand[i] > tar; i-- {
		}
		if i < 0 {
			return [][]int{}
		}
		ret := f(cand[:i], tar)
		for _, item := range f(cand[:i+1], tar-cand[i]) {
			tmp := make([]int, len(item))
			copy(tmp, item)
			ret = append(ret, append(tmp, cand[i]))
		}
		/*
			if _, visited := dp[[2]int{len(cand), tar}]; visited && !equal2(dp[[2]int{len(cand), tar}], ret) {
				fmt.Println(len(cand), tar)
				fmt.Println(dp[[2]int{len(cand), tar}])
				fmt.Println(ret)
				fmt.Println()
			}
		*/
		dp[[2]int{len(cand), tar}] = ret
		return ret
	}
	//defer fmt.Println(dp)
	return f(candidates, target)
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	//fmt.Println(combinationSum([]int{2}, 1))
	//fmt.Println(combinationSum([]int{1}, 1))
	fmt.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}

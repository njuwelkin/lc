package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	ret := [][]int{{}}
	for num, count := range m {
		crtLevel := len(ret)
		tmp := make([]int, count)
		for i := 0; i < count; i++ {
			tmp[i] = num
			for j := 0; j < crtLevel; j++ {
				item := make([]int, len(ret[j])) //, len(ret[j])+i+1)
				copy(item, ret[j])
				ret = append(ret, append(item, tmp[:i+1]...))
			}
		}
		//fmt.Println(ret[:crtLevel])
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{9, 0, 3, 5, 7}))
}

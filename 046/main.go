package main

import "fmt"

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	ret := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		next := [][]int{}
		for _, p := range ret {
			p = append(p, num)
			next = append(next, p)
			for i := 0; i < len(p)-1; i++ {
				tmp := make([]int, len(p))
				copy(tmp, p)

				tmp[i], tmp[len(tmp)-1] = num, tmp[i]
				next = append(next, tmp)
			}
		}
		ret = next
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}

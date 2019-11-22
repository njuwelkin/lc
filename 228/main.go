package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ret := []string{}
	var i, j int
	for i, j = 0, 1; j < len(nums); j++ {
		if nums[j]-nums[j-1] != 1 {
			if j-1 == i {
				ret = append(ret, fmt.Sprintf("%d", nums[i]))
			} else {
				ret = append(ret, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
			}
			i = j
		}
	}
	if j-1 == i {
		ret = append(ret, fmt.Sprintf("%d", nums[i]))
	} else {
		ret = append(ret, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

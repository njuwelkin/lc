package main

import "fmt"

/*
12345
12354
12435
12453
12534
12543
13245
13254
13425
13452
13524
13542
*/
func nextPermutation(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	// find the rightest increase sequence p1, p2
	p1 := -1
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			p1 = i
			break
		}
	}
	p2 := p1 + 1

	// if found, exchange p1 and the smallest value that larger than p1
	if p2 != 0 {
		for j := l - 1; j > p2; j-- {
			if nums[j] > nums[p1] && nums[j] < nums[p2] {
				p2 = j
			}
		}
		nums[p1], nums[p2] = nums[p2], nums[p1]
	}

	// reverse the array of the right of p1
	for i, j := p1+1, l-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	fmt.Println("vim-go")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	for i := 0; i < 20; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}
}

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	ret := 0
	crt := 0
	for _, num := range nums {
		crt += num
		if crt > ret {
			ret = crt
		} else if crt < 0 {
			crt = 0
		}
	}
	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))

	a, b := 1, 2
	fmt.Println(a, b)
	if true {
		a, b := 3, 4
		fmt.Println(a, b)
	}
	fmt.Println(a, b)

	candidate := 5.0
	for i := 1; i < 100; i++ {
		var f float64 = float64(i)
		//fmt.Println(1 - math.Pow(f/(f+1), f))
		t := (f-1)*(candidate-1) + candidate
		fmt.Println(1 - math.Pow((t-1)/t, f))
	}

}

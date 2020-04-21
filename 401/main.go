package main

import "fmt"

func getVals(n int, m int, max int) []int {
	ret := []int{}
	var c func(int, int, int)
	c = func(n, m int, val int) {
		if val >= max {
			return
		}
		if m == 0 {
			ret = append(ret, val)
			return
		}
		c(n-1, m-1, val+(1<<(n-1)))
		if n > m {
			c(n-1, m, val)
		}
	}
	c(n, m, 0)
	return ret
}

func readBinaryWatch(num int) []string {
	ret := []string{}
	for h := 0; h <= num && h < 4; h++ {
		m := num - h
		if m >= 6 {
			continue
		}
		hArray := getVals(4, h, 12)
		mArray := getVals(6, m, 60)
		for _, hVal := range hArray {
			for _, mVal := range mArray {
				if mVal < 10 {
					ret = append(ret, fmt.Sprintf("%d:0%d", hVal, mVal))
				} else {
					ret = append(ret, fmt.Sprintf("%d:%d", hVal, mVal))
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(getVals(4, 3, 12))
	fmt.Println(getVals(4, 0, 12))
	fmt.Println(readBinaryWatch(1))
}

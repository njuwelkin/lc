package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] > pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
	crtTop := 0
	for i := 1; i < len(pairs); i++ {
		if pairs[i][1] <= pairs[crtTop][1] {
			pairs[crtTop] = pairs[i]
		} else if pairs[i][0] <= pairs[crtTop][1] {
			continue
		} else {
			crtTop++
			pairs[crtTop] = pairs[i]
		}
	}
	return crtTop + 1
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
}

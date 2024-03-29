package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	n := len(people)
	ret := make([][]int, n)

	var i, j, count int
	for i = 0; i < n; i++ {
		for j, count = 0, 0; j < n; j++ {
			if ret[j] == nil {
				if count == people[i][1] {
					ret[j] = people[i]
					break
				}
				count++
			}
		}
	}
	return ret
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] == people[b][0] {
			return people[a][1] < people[b][1]
		}
		return people[a][0] > people[b][0]
	})
	fmt.Println(people)

	ret := make([][]int, len(people))
	for _, v := range people {
		copy(ret[v[1]+1:len(ret)], ret[v[1]:len(ret)-1])
		ret[v[1]] = v
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}

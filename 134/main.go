package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	start, end := 0, 0
	tank := 0
	n := len(gas)
	for start < n {
		if tank+gas[end] >= cost[end] {
			tank += gas[end] - cost[end]
			end = (end + 1) % n
			if end == start {
				return start
			}
		} else {
			if start == end {
				start++
				end = start
			} else {
				tank -= gas[start] - cost[start]
				start++
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

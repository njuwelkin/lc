package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMinMoves(machines []int) int {
	sum := 0
	for i := range machines {
		sum += machines[i]
	}
	if sum%len(machines) != 0 {
		return -1
	}
	avg := sum / len(machines)

	count := 0
	leftTarget := 0
	rightTarget := sum - avg
	leftSum := 0
	rightSum := sum
	for i := range machines {
		rightSum -= machines[i]
		toLeft, toRight := 0, 0
		if leftSum < leftTarget {
			toLeft = leftTarget - leftSum
		}
		if rightSum < rightTarget {
			toRight = rightTarget - rightSum
		}
		leftSum += machines[i]
		rightTarget -= avg
		leftTarget += avg
		count = max(count, toLeft+toRight)
	}

	return count
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findMinMoves([]int{1, 0, 5}))
	fmt.Println(findMinMoves([]int{0, 3, 0}))
	fmt.Println(findMinMoves([]int{0, 2, 0}))
	fmt.Println(findMinMoves([]int{4, 0, 0, 4}))
}

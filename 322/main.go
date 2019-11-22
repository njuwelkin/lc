package main

import "fmt"

func coinChange(coins []int, amount int) int {
	q := []int{amount}
	visited := map[int]bool{}
	level := 0
	for len(q) != 0 {
		level++
		fmt.Println(q)
		next := []int{}
		for _, num := range q {
			for _, c := range coins {
				if c == num {
					return level
				} else if c > num {
					continue
				}
				if _, ok := visited[num-c]; ok {
					continue
				}
				visited[num-c] = true
				next = append(next, num-c)
			}
		}
		q = next
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{1, 2, 5}, 100))
}

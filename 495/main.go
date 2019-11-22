package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	} else if n == 1 {
		return duration
	}
	ret := 0
	var i, j int
	for i, j = 0, 1; j < n; j++ {
		if timeSeries[j]-timeSeries[i] > duration {
			ret += timeSeries[j-1] - timeSeries[i] + duration
			i = j
		}
	}
	ret += timeSeries[j-1] - timeSeries[i] + duration
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}

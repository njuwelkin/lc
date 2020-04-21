package main

import "fmt"

func findPoisonedDuration2(timeSeries []int, duration int) int {
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

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	ret := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] < duration {
			ret += timeSeries[i] - timeSeries[i-1]
		} else {
			ret += duration
		}
	}
	ret += duration
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration2([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
	fmt.Println(findPoisonedDuration2([]int{1, 2}, 2))
}

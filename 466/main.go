package main

import "fmt"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	var i, j, idx int
	startIdx := -1
	for i = 0; i < n1*len(s1); i++ {
		idx = i % len(s1)
		if s1[idx] == s2[j] {
			j++
			if j == len(s2) {
				i++
				startIdx = i
				break
			}
		}
	}
	if startIdx == -1 {
		return 0
	}
	//fmt.Println("startIdx", startIdx)

	endIdx := -1
	count := 0
	for j = 0; i < n1*len(s1); i++ {
		idx = i % len(s1)
		if s1[idx] == s2[j] {
			j++
			if j == len(s2) {
				j = 0
				count++
				if (i+1)%len(s1) == startIdx%len(s1) {
					i++
					endIdx = i
					break
				}
			}
		}
	}
	//fmt.Println("count", count)
	if endIdx == -1 {
		return (count + 1) / n2
	}
	//fmt.Println("endIdx", endIdx)
	length := endIdx - startIdx
	count = count*((n1*len(s1)-startIdx)/length) + 1

	j = 0
	for i = n1*len(s1) - ((n1*len(s1) - startIdx) % length); i < n1*len(s1); i++ {
		idx = i % len(s1)
		if s1[idx] == s2[j] {
			j++
			if j == len(s2) {
				j = 0
				count++
			}
		}
	}

	return count / n2
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(getMaxRepetitions("acb", 4, "ab", 2))
	fmt.Println(getMaxRepetitions("aaa", 3, "aa", 1))
}

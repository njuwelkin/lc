package main

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	if input == "" {
		return 0
	}
	lengthArray := []int{}
	paths := strings.Split(input, "\n")
	ret := 0
	for _, path := range paths {
		var i int
		for i = 0; i < len(path) && path[i] == '\t'; i++ {
		}
		if i == len(path) {
			return -1
		}
		for i >= len(lengthArray) {
			lengthArray = append(lengthArray, 0)
		}
		if i == 0 {
			lengthArray[i] = len(path[i:])
		} else {
			lengthArray[i] = lengthArray[i-1] + len(path[i:]) + 1
		}
		if lengthArray[i] > ret {
			ret = lengthArray[i]
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(lengthLongestPath(""))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}

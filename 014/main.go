package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ret := strs[0]
	for _, str := range strs[1:] {
		var i int
		for i = 0; i < min(len(ret), len(str)); i++ {
			if ret[i] != str[i] {
				break
			}
		}
		ret = ret[:i]
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

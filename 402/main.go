package main

import "fmt"

func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}

	lenRet := n - k
	ret := make([]byte, lenRet) // it's a stack

	var i, j int
	for i, j = 0, 0; j < n && i < lenRet; {
		if i == 0 || num[j] > ret[i-1] || k == 0 {
			if !(i == 0 && num[j] == '0') {
				ret[i] = num[j] //push
				i++
			} else {
				lenRet--
			}
			j++
		} else {
			k--
			i-- //pop
		}
	}
	if i == 0 {
		return "0"
	}
	return string(ret[:i])
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
}

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

func removeKdigits2(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	stack := make([]byte, len(num))
	top := 0

	var i int
	for i = 0; i < len(num); {
		if top == 0 || num[i] >= stack[top-1] || k == 0 {
			stack[top] = num[i] //push
			top++
			i++
		} else {
			top-- // pop
			k--
		}
	}

	stack = stack[:top-k]
	for i = 0; i < len(stack) && stack[i] == '0'; i++ {
	}
	if i == len(stack) {
		return "0"
	}
	return string(stack[i:])
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
}

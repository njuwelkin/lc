package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := make([]int, len(s)+1)
	stack[0] = -1
	top := 1

	ret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack[top] = i
			top++
		} else {
			if top > 1 {
				top--
				if i-stack[top-1] > ret {
					ret = i - stack[top-1]
				}
			} else {
				stack[top-1] = i
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}

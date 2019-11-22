package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := 0
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err != nil {
			if top < 2 {
				return 1 << 62
			}
			top--
			switch token {
			case "+":
				stack[top-1] += stack[top]
			case "-":
				stack[top-1] -= stack[top]
			case "*":
				stack[top-1] *= stack[top]
			case "/":
				stack[top-1] /= stack[top]
			default:
				return 1 << 62
			}
		} else {
			stack[top] = val
			top++
		}
	}
	if top != 1 {
		return 1 << 62
	}
	return stack[0]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}

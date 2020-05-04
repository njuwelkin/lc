package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := 0
	push := func(c byte) {
		stack[top] = c
		top++
	}
	pop := func() byte {
		if top == 0 {
			return 0
		}
		top--
		return stack[top]
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			push(s[i])
		} else {
			var expect byte
			switch s[i] {
			case ')':
				expect = '('
			case ']':
				expect = '['
			case '}':
				expect = '{'
			default:
				return false
			}
			if pop() != expect {
				return false
			}
		}
	}
	return top == 0
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

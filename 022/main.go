package main

import "fmt"

func generateParenthesis(n int) []string {
	buf := make([]byte, n*2)
	ret := []string{}

	var gen func(int, int, int)
	gen = func(idx int, left, right int) {
		if idx == n*2 {
			ret = append(ret, string(buf))
			return
		}
		if left < n {
			buf[idx] = '('
			gen(idx+1, left+1, right)
		}
		if right < left {
			buf[idx] = ')'
			gen(idx+1, left, right+1)
		}

	}

	gen(0, 0, 0)
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}

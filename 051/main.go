package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createPatterns(n int) []string {
	ret := []string{}
	pattern := make([]byte, n)
	for i := 0; i < n; i++ {
		pattern[i] = '.'
	}
	for i := 0; i < n; i++ {
		pattern[i] = 'Q'
		if i != 0 {
			pattern[i-1] = '.'
		}
		ret = append(ret, string(pattern))
	}
	return ret
}

func symmetry(solutions [][]string) [][]string {
	n := len(solutions)
	for i := 0; i < n; i++ {
		l := len(solutions[i])
		s := make([]string, l)
		for j := 0; j < len(s); j++ {
			s[j] = solutions[i][l-1-j]
		}
		solutions = append(solutions, s)
	}
	return solutions
}

func solveNQueens(n int) [][]string {
	ret := [][]string{}
	patterns := createPatterns(n)

	rows := make([]int, n)
	var f func(line int)
	f = func(line int) {
		if line == n {
			solution := []string{}
			for i := 0; i < n; i++ {
				solution = append(solution, patterns[rows[i]])
			}
			ret = append(ret, solution)
			return
		}
		for i := 0; i < n; i++ {
			conflict := false
			for j := 0; j < line; j++ {
				if rows[j] == i || abs(i-rows[j]) == line-j {
					conflict = true
					break
				}
			}
			if conflict {
				continue
			}
			rows[line] = i
			f(line + 1)
		}
	}
	for i := 0; i < (n+1)/2; i++ {
		rows[0] = i
		f(1)
	}
	return symmetry(ret)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(5))
}

package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return word == ""
	}
	n := len(board[0])

	var search(int, int, int) bool
	search = func(i, j, wordIdx int) bool {
		if wordIdx == len(word) {
			return true
		}
		if i >= m || i < 0 ||
			j >= n || j < 0 ||
			board[i][j] != word[wordIdx] {
			return false
		}
		for _, delta := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			if search(i+delta[0], j+delta[1], wordIdx+1) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
}

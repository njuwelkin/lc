package main

import "fmt"

func printBoard(board [][]byte) {
	for i := range board {
		for j := range board[0] {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	forAllNeibor := func(x int, y int, f func(i, j int)) {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(board) {
				continue
			}
			for j := y - 1; j <= y+1; j++ {
				if j < 0 || j >= len(board[i]) {
					continue
				}
				if !(i == x && j == y) {
					f(i, j)
				}
			}
		}

	}

	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	q := [][]int{[]int{x, y}}
	for len(q) > 0 {
		next := [][]int{}
		for i := range q {
			x, y = q[i][0], q[i][1]
			var count byte = 0
			forAllNeibor(x, y, func(i, j int) {
				if board[i][j] == 'M' {
					count++
				}
			})
			if count != 0 {
				board[x][y] = '0' + count
			} else {
				board[x][y] = 'B'
				forAllNeibor(x, y, func(i, j int) {
					if board[i][j] == 'E' {
						board[i][j] = 'e'
						next = append(next, []int{i, j})
					}

				})
			}
		}
		q = next
	}
	return board
}

func main() {
	fmt.Println("vim-go")
	printBoard(updateBoard([][]byte{{'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'M', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}, {'E', 'E', 'E', 'E', 'E'}}, []int{3, 0}))
}

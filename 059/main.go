package main

import "fmt"

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	crtVal := 1
	for i := 0; i < n/2; i++ {
		m := n - i
		zero := i
		for j := zero; j < m-1; j++ {
			ret[zero][j] = crtVal
			crtVal++
		}
		for j := zero; j < m-1; j++ {
			ret[j][m-1] = crtVal
			crtVal++
		}
		for j := m - 1; j >= zero+1; j-- {
			ret[m-1][j] = crtVal
			crtVal++
		}
		for j := m - 1; j >= zero+1; j-- {
			ret[j][zero] = crtVal
			crtVal++
		}
	}
	if n%2 != 0 {
		ret[n/2][n/2] = crtVal
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(generateMatrix(0))
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}

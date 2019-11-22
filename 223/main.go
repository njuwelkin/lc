package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	overlapX := 0
	if max(A, E) < min(C, G) {
		overlapX = min(C, G) - max(A, E)
	}
	overlapY := 0
	if max(B, F) < min(D, H) {
		overlapY = min(D, H) - max(B, F)
	}
	fmt.Println(overlapX, overlapY)
	return (C-A)*(D-B) + (G-E)*(H-F) - overlapX*overlapY
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}

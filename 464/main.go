package main

import "fmt"

func bf(maxChoosableInteger int, desiredTotal int) bool {
	fmt.Println("desiredTotal: ", desiredTotal)
	visited := make([]bool, maxChoosableInteger+1)
	path := make([]int, 100)
	tmp := -1
	var f func(int, int) bool
	f = func(des int, depth int) bool {
		if des <= 0 {
			return false
		}
		//fmt.Println("depth:", depth, des)
		for i := 1; i <= maxChoosableInteger; i++ {
			if !visited[i] {
				path[depth] = i
				visited[i] = true
				res := f(des-i, depth+1)
				visited[i] = false
				if !res {
					//fmt.Println(i)
					tmp = i
					return true
				}
			}
		}
		return false
	}
	res := f(desiredTotal, 0)
	fmt.Println(tmp)
	return res
}

func setBit(val *int, pos uint) int {
	*val |= 1 << pos
	return *val
}

func clearBit(val *int, pos uint) int {
	*val &= ^(1 << pos)
	return *val
}

func getBit(val int, pos uint) bool {
	return (val & (1 << pos)) != 0
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	fmt.Println("desiredTotal: ", desiredTotal)
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	var hit int = 0

	var visited int
	dp := map[[2]int]bool{}
	var f func(int) bool
	f = func(des int) bool {
		if des <= 0 {
			return false
		}
		if ret, found := dp[[2]int{visited, des}]; found {
			hit++
			return ret
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			//for i := maxChoosableInteger; i >= 1; i-- {
			if !getBit(visited, uint(i)) {
				setBit(&visited, uint(i))
				res := f(des - i)
				clearBit(&visited, uint(i))
				if !res {
					//fmt.Println(i)
					dp[[2]int{visited, des}] = true
					return true
				}
			}
		}
		dp[[2]int{visited, des}] = false
		return false
	}
	f(desiredTotal)
	fmt.Println(len(dp), hit)
	return f(desiredTotal)
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(canIWin(9, 10))
	for i := 100; i < 200; i++ {
		fmt.Println(canIWin(20, i))
	}
}

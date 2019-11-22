package main

import (
	"fmt"
	"sort"
)

func twoSumCount(A, B []int, target int) int {
	n := len(A)
	ret := 0
	for i, j := 0, n-1; ; {
		sum := A[i] + B[j]
		if sum == target {
			countI, countJ := 1, 1
			var k int
			for k = i + 1; k < n && A[k] == A[i]; k++ {
				countI++
			}
			i = k
			for k = j - 1; k >= 0 && B[k] == B[j]; k-- {
				countJ++
			}
			j = k
			ret += countI * countJ
		} else if sum < target {
			i++
		} else {
			j--
		}
		if i == n || j < 0 {
			return ret
		}
	}
	return ret
}

func fourSumCount1(A []int, B []int, C []int, D []int) int {
	n := len(A)
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	sort.Ints(D)
	fmt.Println(A, B, C, D)

	ret := 0
	for i := 0; i < n; {
		retI := 0
		for j := 0; j < n; {
			target := -(A[i] + B[j])
			retJ := twoSumCount(C, D, target)
			k := j
			for j = j + 1; j < n && B[j] == B[k]; j++ {
			}
			retI += (j - k) * retJ
		}
		k := i
		for i = i + 1; i < n && A[i] == A[k]; i++ {
		}
		ret += (i - k) * retI
	}
	return ret
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := map[int]int{}
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if _, ok := m[A[i]+B[j]]; !ok {
				m[A[i]+B[j]] = 1
			} else {
				m[A[i]+B[j]]++
			}
		}
	}

	ret := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if count, ok := m[-C[i]-D[j]]; ok {
				ret += count
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	//fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	fmt.Println(fourSumCount([]int{0, 1, -1}, []int{-1, 1, 0}, []int{0, 0, 1}, []int{-1, 1, 1}))
}

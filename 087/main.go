package main

import "fmt"

func isScramble1(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var m = map[byte]int{}
	for i, _ := range s1 {
		m[s1[i]]++
		m[s2[i]]--
	}

	for _, j := range m {
		if j != 0 {
			return false
		}
	}

	n := len(s1)
	if n == 1 {
		return s1 == s2
	}
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) ||
			isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}

func calcWeight(s string) []int {
	ret := make([]int, len(s)+1)
	ret[0] = 0
	for i := 1; i < len(ret); i++ {
		ret[i] = ret[i-1] + int(s[i-1])
	}
	return ret
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	w1 := calcWeight(s1)
	w2 := calcWeight(s2)

	dp := map[[4]int]bool{}
	var f func(i1, j1, i2, j2 int) bool
	f = func(i1, j1, i2, j2 int) bool {
		if ret, ok := dp[[4]int{i1, j1, i2, j2}]; ok {
			return ret
		}
		if s1[i1:j1] == s2[i2:j2] {
			dp[[4]int{i1, j1, i2, j2}] = true
			return true
		}
		if w1[j1]-w1[i1] != w2[j2]-w2[i2] {
			dp[[4]int{i1, j1, i2, j2}] = false
			return false
		}
		for k := 1; k < j1-i1; k++ {
			if f(i1, i1+k, i2, i2+k) && f(i1+k, j1, i2+k, j2) ||
				f(i1, i1+k, j2-k, j2) && f(i1+k, j1, i2, i2+(j1-i1-k)) {
				dp[[4]int{i1, j1, i2, j2}] = true
				return true
			}
		}
		dp[[4]int{i1, j1, i2, j2}] = false
		return false
	}
	return f(0, len(s1), 0, len(s2))
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
}

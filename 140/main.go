package main

import "fmt"

type TrieNode struct {
	Next     [26]*TrieNode
	EndPoint bool
}

func (dict *TrieNode) Insert(s string) {
	p := dict
	for i := 0; i < len(s); i++ {
		if p.Next[s[i]-'a'] == nil {
			p.Next[s[i]-'a'] = new(TrieNode)
		}
		p = p.Next[s[i]-'a']
	}
	p.EndPoint = true
}

func wordBreak(s string, wordDict []string) []string {
	dict := TrieNode{}
	for _, w := range wordDict {
		dict.Insert(w)
	}

	dp := make([][]string, len(s))
	visited := make([]bool, len(s))
	var f func(int) []string
	f = func(idx int) []string {
		if idx == len(s) {
			return []string{""}
		}
		if visited[idx] {
			return dp[idx]
		}
		visited[idx] = true
		dp[idx] = []string{}
		for i, p := idx, &dict; i < len(s) && p != nil; i++ {
			p = p.Next[s[i]-'a']
			if p != nil && p.EndPoint {
				for _, str := range f(i + 1) {
					if str == "" {
						dp[idx] = append(dp[idx], s[idx:i+1])
					} else {
						dp[idx] = append(dp[idx], fmt.Sprintf("%s %s", s[idx:i+1], str))
					}
				}
			}
		}

		return dp[idx]
	}
	return f(0)
}

func wordBreak1(s string, wordDict []string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}

	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	canBreak := make([][]int, n)
	for i := len(s) - 1; i >= 0; i-- {
		canBreak[i] = []int{}
		if m[s[i:]] {
			canBreak[i] = append(canBreak[i], n)
		}
		for j := i + 1; j < n; j++ {
			if len(canBreak[j]) != 0 && m[s[i:j]] {
				canBreak[i] = append(canBreak[i], j)
			}
		}
	}
	fmt.Println(canBreak)

	ret := []string{}
	var dfs func(int, string)
	dfs = func(idx int, str string) {
		if idx == n {
			ret = append(ret, str[:len(str)-1])
			return
		}
		for _, next := range canBreak[idx] {
			dfs(next, str+s[idx:next]+" ")
		}
	}
	dfs(0, "")
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(wordBreak("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}))
}

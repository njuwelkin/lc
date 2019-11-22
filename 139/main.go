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

func wordBreak1(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	canBreak := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if m[s[:i+1]] {
			canBreak[i] = true
		} else {
			for j := 0; j < i; j++ {
				if canBreak[j] && m[s[j+1:i+1]] {
					canBreak[i] = true
					break
				}
			}
		}
	}
	fmt.Println(canBreak)
	return canBreak[len(s)-1]
}

func wordBreak(s string, wordDict []string) bool {
	dict := TrieNode{}
	for _, w := range wordDict {
		dict.Insert(w)
	}

	dp := make([]bool, len(s))
	visited := make([]bool, len(s))
	var f func(int) bool
	f = func(idx int) bool {
		if idx == len(s) {
			return true
		}
		if visited[idx] {
			return dp[idx]
		}
		visited[idx] = true
		for i, p := idx, &dict; i < len(s) && p != nil; i++ {
			p = p.Next[s[i]-'a']
			if p != nil && p.EndPoint {
				if f(i + 1) {
					dp[idx] = true
					return true
				}
			}
		}

		dp[idx] = false
		return false
	}
	return f(0)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

package main

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	Next     [26]*TrieNode
	EndPoint bool
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dict := TrieNode{}
	ret := []string{}

	var findNext func(string) bool
	findNext = func(w string) bool {
		node := &dict
		var i int
		for i = 0; i < len(w); i++ {
			if node.EndPoint {
				if findNext(w[i:]) {
					return true
				}
			}
			if node.Next[w[i]-'a'] == nil {
				return false
			}
			node = node.Next[w[i]-'a']
		}
		return node.EndPoint
	}

	add := func(w string) {
		node := &dict
		var i int
		for i = 0; i < len(w); i++ {
			if node.EndPoint {
				if findNext(w[i:]) {
					ret = append(ret, w)
					return
				}
			}
			if node.Next[w[i]-'a'] != nil {
				node = node.Next[w[i]-'a']
			} else {
				break
			}
		}
		if i == len(w) {
			return
		}
		for ; i < len(w); i++ {
			node.Next[w[i]-'a'] = new(TrieNode)
			node = node.Next[w[i]-'a']
		}
		node.EndPoint = true
	}

	for _, word := range words {
		add(word)
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}

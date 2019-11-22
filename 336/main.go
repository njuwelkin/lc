package main

import (
	"fmt"
	"sort"
)

type DictNode struct {
	next [26]*DictNode
	idx  int
}

func (d *DictNode) Insert(word string, idx int) {
	p := d
	for i := 0; i < len(word); i++ {
		if p.next[word[i]-'a'] == nil {
			p.next[word[i]-'a'] = new(DictNode)
			p.next[word[i]-'a'].idx = -1
		}
		p = p.next[word[i]-'a']
	}
	p.idx = idx
}

func (d *DictNode) FindPair(w string, includingEnd bool) []int {
	p := d
	ret := []int{}
	for i := 0; i < len(w) && p != nil; i++ {
		if p.idx != -1 {
			if isPal(w[i:]) {
				ret = append(ret, p.idx)
			}
		}
		p = p.next[w[i]-'a']
	}
	if p != nil && p.idx != -1 && includingEnd {
		ret = append(ret, p.idx)
	}
	return ret
}

func isPal(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		ret[n-1-i] = s[i]
	}
	return string(ret)
}

func palindromePairs(words []string) [][]int {
	idx := make([]int, len(words))
	for i, _ := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return len(words[idx[j]]) < len(words[idx[i]])
	})
	//fmt.Println(idx)

	d1 := DictNode{}
	d2 := DictNode{}
	d1.idx, d2.idx = -1, -1
	for i, w := range words {
		d1.Insert(w, i)
		d2.Insert(reverse(w), i)
	}

	ret := [][]int{}
	for _, i := range idx {
		w := words[i]
		tmp := d1.FindPair(reverse(w), true)
		for _, left := range tmp {
			if left != i {
				ret = append(ret, []int{left, i})
			}
		}
		tmp = d2.FindPair(w, false)
		for _, right := range tmp {
			if right != i {
				ret = append(ret, []int{i, right})
			}
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(palindromePairs([]string{"a", ""}))
	//fmt.Println(palindromePairs([]string{"ls", "l", "s"}))
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
}

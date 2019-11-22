package main

import (
	"fmt"
	"sort"
)

type node struct {
	c     byte
	count int
}

func frequencySort(s string) string {
	count := make([]int, 256)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	buf := make([]node, 256)
	for i := 0; i < 256; i++ {
		buf[i] = node{byte(i), count[i]}
	}

	sort.Slice(buf, func(i, j int) bool {
		return buf[i].count > buf[j].count
	})

	ret := []byte{}
	for i := 0; i < 256; i++ {
		for j := 0; j < buf[i].count; j++ {
			ret = append(ret, buf[i].c)
		}
	}
	return string(ret)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(frequencySort("Aabb"))
}

package main

import (
	"bytes"
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	ret := []string{}

	i := 0
	for i < len(words) {
		var buf bytes.Buffer
		buf.Reset()
		buf.WriteString(words[i])

		var j int
		remain := maxWidth
		remain -= len(words[i])
		for j = i + 1; j < len(words) && remain >= len(words[j])+1; j++ {
			remain -= len(words[j]) + 1
		}

		//fmt.Println(words[i:j])
		if j-i == 1 {
			i = j
			for k := 0; k < remain; k++ {
				buf.WriteByte(' ')
			}
			ret = append(ret, buf.String())
			continue
		}

		countSlot := j - i - 1
		totalSpace := remain + countSlot
		spaceWidth := totalSpace / countSlot

		for i = i + 1; i < j; i++ {
			for k := 0; k < spaceWidth; k++ {
				buf.WriteByte(' ')
			}
			if totalSpace > countSlot*spaceWidth {
				buf.WriteByte(' ')
				totalSpace -= 1
			}
			totalSpace -= spaceWidth
			countSlot -= 1
			buf.WriteString(words[i])
		}
		ret = append(ret, buf.String())
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	ret := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for _, line := range ret {
		fmt.Println(line)
	}
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
}

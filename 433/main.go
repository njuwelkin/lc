package main

import "fmt"

func encode(s string) int {
	dict := [26]int{}
	dict['A'-'A'] = 0
	dict['C'-'A'] = 1
	dict['G'-'A'] = 2
	dict['T'-'A'] = 3
	ret := 0
	for i := 0; i < len(s); i++ {
		ret = ret<<2 + dict[s[i]-'A']
	}
	return ret
}

func isNeibor(g1, g2 int) bool {
	mask := 3
	count := 0
	g := g1 ^ g2
	for i := 0; i < 8; i++ {
		if mask&g != 0 {
			count++
			if count > 1 {
				return false
			}
		}
		mask <<= 2
	}
	return true
}

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	// locate start and end in bank
	endIdx := -1
	for i, item := range bank {
		if item == end {
			endIdx = i
			break
		}
	}
	if endIdx < 0 {
		bank = append(bank, end)
		endIdx = len(bank) - 1
	}
	bank = append(bank, start)
	startIdx := len(bank) - 1
	//fmt.Println(bank)
	//fmt.Println(startIdx, endIdx)

	// encode
	encodeBank := make([]int, len(bank))
	for i, item := range bank {
		encodeBank[i] = encode(item)
	}
	//fmt.Println(encodeBank)

	// bfs
	visited := make([]bool, len(bank))
	crtLevel := []int{startIdx}
	visited[startIdx] = true
	depth := 1
	for len(crtLevel) != 0 {
		//fmt.Println(crtLevel)
		nextLeval := []int{}
		for _, node := range crtLevel {
			orig := encodeBank[node]
			for i, target := range encodeBank {
				if !visited[i] && isNeibor(orig, target) {
					if i == endIdx {
						return depth
					}
					nextLeval = append(nextLeval, i)
					visited[i] = true
				}
			}
		}
		crtLevel = nextLeval
		depth++
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
}

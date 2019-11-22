package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	maxLen := 0
	tailNode := -1
	mid1, mid2 := 0, 0
	path := make([]int, n)
	visited := make([]bool, n)
	var maxDepth func(root int, depth int)
	maxDepth = func(root int, depth int) {
		path[depth] = root
		if depth > maxLen {
			maxLen = depth
			tailNode = root
			mid2 = path[(depth+1)/2]
			if (depth+1)%2 == 0 {
				mid1 = path[(depth+1)/2-1]
			} else {
				mid1 = mid2
			}
		}
		for _, sub := range graph[root] {
			if !visited[sub] {
				visited[sub] = true
				maxDepth(sub, depth+1)
			}
		}
	}

	visited[0] = true
	maxDepth(0, 0)

	visited = make([]bool, n)
	visited[tailNode] = true
	maxDepth(tailNode, 0)

	if mid1 == mid2 {
		return []int{mid2}
	}
	return []int{mid1, mid2}
}
func main() {
	fmt.Println("vim-go")
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}))
}

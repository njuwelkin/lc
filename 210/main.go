package main

import "fmt"

type Node struct {
	countPrev  int
	seccessors []int
}

//func canFinish(numCourses int, prerequisites [][]int) bool {
func findOrder(numCourses int, prerequisites [][]int) []int {
	courses := make([]Node, numCourses)
	for i, _ := range courses {
		courses[i].seccessors = []int{}
	}
	for _, pre := range prerequisites {
		courses[pre[1]].seccessors = append(courses[pre[1]].seccessors, pre[0])
		courses[pre[0]].countPrev++
	}

	queue := []int{}
	for i, course := range courses {
		if course.countPrev == 0 {
			queue = append(queue, i)
		}
	}

	ret := make([]int, numCourses)
	idx := 0
	for len(queue) != 0 {
		ret[idx] = queue[0]
		idx++
		course := courses[queue[0]]
		queue = queue[1:]
		for _, secc := range course.seccessors {
			courses[secc].countPrev--
			if courses[secc].countPrev == 0 {
				queue = append(queue, secc)
			}
		}
	}

	if idx != numCourses {
		return []int{}
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

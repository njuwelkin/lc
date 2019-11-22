package main

import "fmt"

type Node struct {
	countPrev  int
	seccessors []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([]Node, numCourses)
	for i, _ := range courses {
		courses[i].seccessors = []int{}
	}
	for _, pre := range prerequisites {
		courses[pre[0]].seccessors = append(courses[pre[0]].seccessors, pre[1])
		courses[pre[1]].countPrev++
	}

	queue := []int{}
	for i, course := range courses {
		if course.countPrev == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		course := courses[queue[0]]
		queue = queue[1:]
		for _, secc := range course.seccessors {
			courses[secc].countPrev--
			if courses[secc].countPrev == 0 {
				queue = append(queue, secc)
			}
		}
	}

	for _, course := range courses {
		if course.countPrev != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

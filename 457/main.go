package main

import "fmt"

func hasLoop(nums []int, visited []bool, head int) bool {
	n := len(nums)
	if visited[head] {
		return false
	}
	visited[head] = true
	isPos := 1
	if nums[head] < 0 {
		isPos = -1
	}
	fast, slow := head, head
	for {
		fast = (fast + nums[fast]%n + n) % n
		if nums[fast]*isPos < 0 {
			return false
		}
		visited[fast] = true
		if fast == slow {
			break
		}

		fast = (fast + nums[fast]%n + n) % n
		if nums[fast]*isPos < 0 {
			return false
		}
		visited[fast] = true
		if fast == slow {
			break
		}
		slow = (slow + nums[slow]%n + n) % n
	}
	return (fast+nums[fast]+n)%n != fast
}

func circularArrayLoop(nums []int) bool {
	visited := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if hasLoop(nums, visited, i) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{-1, 2}))
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
}

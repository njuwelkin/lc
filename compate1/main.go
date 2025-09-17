package main

import "fmt"

/*
a和b轮流在桌子上从左到右写下1或0，a先手，当写到8位时需要擦掉第一位，这样桌子上始终有一个7位二进制数
记录所有出现过的7位二进制数，当某人写下1或0擦掉第一位后与之前记录过二进制数出现重复则获胜
*/

// DPMemoKey 定义dp映射的键类型，包含state和访问状态信息
type DPMemoKey struct {
	state    uint16 // 7位二进制状态
	visited0 uint64 // 访问状态0-63
	visited1 uint64 // 访问状态64-127
}

// VisitedTracker 封装两个64位整数用于跟踪访问状态
type VisitedTracker struct {
	visited0 uint64 // 存储0-63状态
	visited1 uint64 // 存储64-127状态
}

// IsVisited 检查指定状态是否被访问过
func (vt *VisitedTracker) IsVisited(state int) bool {
	if state < 64 {
		return (vt.visited0 & (1 << state)) != 0
	} else {
		return (vt.visited1 & (1 << (state - 64))) != 0
	}
}

// SetVisited 标记指定状态为已访问
func (vt *VisitedTracker) SetVisited(state int) {
	if state < 64 {
		vt.visited0 |= (1 << state)
	} else {
		vt.visited1 |= (1 << (state - 64))
	}
}

// UnsetVisited 取消标记指定状态
func (vt *VisitedTracker) UnsetVisited(state int) {
	if state < 64 {
		vt.visited0 &^= (1 << state)
	} else {
		vt.visited1 &^= (1 << (state - 64))
	}
}

// GetVisitedStates 获取当前的访问状态
func (vt *VisitedTracker) GetVisitedStates() (uint64, uint64) {
	return vt.visited0, vt.visited1
}

// DFS_RESULT 定义DFS结果的枚举类型
type DFS_RESULT int

const (
	// P0Win 表示先手获胜
	P0Win DFS_RESULT = iota
	// P1Win 表示后手获胜
	P1Win
)

func solution() {
	// 创建VisitedTracker实例替代原来的两个变量
	visitedTracker := &VisitedTracker{0, 0}
	// 修改dp的key为DPMemoKey类型，包含state、visited0和visited1
	dp := make(map[DPMemoKey]bool)
	longestPath := 0
	var dfs func(int, int, int) DFS_RESULT
	// dfs performs a depth-first search to determine if the current player can win from the given state
	// Parameters:
	//   state: The current 7-bit binary state representing the numbers on the table
	//   depth: The current search depth, used to determine which player's turn it is (0 for first player, 1 for second)
	// Returns:
	//   P0Win if the first player can force a win from this state
	//   P1Win if the second player can force a win from this state
	dfs = func(state int, depth int, path int) DFS_RESULT {
		if path > longestPath {
			longestPath = path
		}
		// 获取当前的访问状态
		v0, v1 := visitedTracker.GetVisitedStates()
		// 创建dp的key
		key := DPMemoKey{
			state:    uint16(state),
			visited0: v0,
			visited1: v1,
		}
		// Check memoization table first to avoid redundant calculations
		if win, ok := dp[key]; ok {
			if win {
				return DFS_RESULT(depth % 2)
			} else {
				return DFS_RESULT(1 - depth%2)
			}
		}

		// Calculate new state by shifting left (erasing first bit) and keeping only 7 bits
		newState := (state << 1) & (1<<7 - 1)
		path = path << 1

		// Try both possible moves: writing 0 and writing 1
		for i := 0; i < 2; i++ {
			currentState := newState + i

			// 使用VisitedTracker的方法检查状态是否被访问过
			if visitedTracker.IsVisited(currentState) {
				updateDPStatus(dp, key, true)
				return DFS_RESULT(depth % 2) // i win
			}

			// 使用VisitedTracker的方法标记状态为已访问
			visitedTracker.SetVisited(currentState)
			ret := dfs(newState+i, depth+1, path+i)

			// 使用VisitedTracker的方法取消标记状态（回溯）
			visitedTracker.UnsetVisited(currentState)

			// If this move leads to a win for current player, memoize and return
			if ret == DFS_RESULT(depth%2) { // i finally win
				updateDPStatus(dp, key, true)
				return ret
			}
		}

		// If neither move leads to a win, current player loses
		result := DFS_RESULT(1 - depth%2) // lose
		dp[key] = false
		return result
	}
	win := dfs(0, 0, 0)
	fmt.Printf("longestPath: %b\n", longestPath)
	fmt.Println("win:", win)
}

func updateDPStatus(dp map[DPMemoKey]bool, key DPMemoKey, win bool) {
	fmt.Printf("updateDPStatus: visited: %064b%064b, state: %b, win: %v\n", key.visited1, key.visited0, key.state, win)
	dp[key] = win
}

func main() {
	fmt.Println("vim-go")
	solution()
}

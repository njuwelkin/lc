package main

import "fmt"

/*
a和b轮流在桌子上从左到右写下1或0，a先手，当写到8位时需要擦掉第一位
这样桌子上始终有一个7位二进制数，记录所有出现过的7位二进制数，
当某人写下1或0擦掉第一位后与之前记录过二进制数出现重复则获胜，请问先手或后手有没有必胜策略
*/

// DPMemoKey defines the key type for the dp map, containing state and visited status information
type DPMemoKey struct {
	state    uint16 // 7-bit binary state
	visited0 uint64 // visited status for 0-63
	visited1 uint64 // visited status for 64-127
}

// VisitedTracker encapsulates two 64-bit integers for tracking visited states
type VisitedTracker struct {
	visited0 uint64 // stores states 0-63
	visited1 uint64 // stores states 64-127
}

// IsVisited checks if the specified state has been visited
func (vt *VisitedTracker) IsVisited(state int) bool {
	if state < 64 {
		return (vt.visited0 & (1 << state)) != 0
	} else {
		return (vt.visited1 & (1 << (state - 64))) != 0
	}
}

// SetVisited marks the specified state as visited
func (vt *VisitedTracker) SetVisited(state int) {
	if state < 64 {
		vt.visited0 |= (1 << state)
	} else {
		vt.visited1 |= (1 << (state - 64))
	}
}

// UnsetVisited unmarks the specified state
func (vt *VisitedTracker) UnsetVisited(state int) {
	if state < 64 {
		vt.visited0 &^= (1 << state)
	} else {
		vt.visited1 &^= (1 << (state - 64))
	}
}

// GetVisitedStates gets the current visited states
func (vt *VisitedTracker) GetVisitedStates() (uint64, uint64) {
	return vt.visited0, vt.visited1
}

// DFS_RESULT defines the enumeration type for DFS results
type DFS_RESULT int

const (
	// P0Win indicates the first player wins
	P0Win DFS_RESULT = iota
	// P1Win indicates the second player wins
	P1Win
)

func solution() {
	// Create VisitedTracker instance to replace the original two variables
	visitedTracker := &VisitedTracker{0, 0}
	// Modify dp's key to DPMemoKey type, including state, visited0 and visited1
	dp := make(map[DPMemoKey]bool)
	var dfs func(int, int) DFS_RESULT
	// dfs performs a depth-first search to determine if the current player can win from the given state
	// Parameters:
	//   state: The current 7-bit binary state representing the numbers on the table
	//   depth: The current search depth, used to determine which player's turn it is (0 for first player, 1 for second)
	// Returns:
	//   P0Win if the first player can force a win from this state
	//   P1Win if the second player can force a win from this state
	dfs = func(state int, depth int) DFS_RESULT {
		// Get current visited states
		v0, v1 := visitedTracker.GetVisitedStates()
		// Create key for dp
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

		// Try both possible moves: writing 0 and writing 1
		for i := 0; i < 2; i++ {
			currentState := newState + i

			// Use VisitedTracker method to check if state has been visited
			if visitedTracker.IsVisited(currentState) {
				updateDPStatus(dp, key, true)
				return DFS_RESULT(depth % 2) // i win
			}

			// Use VisitedTracker method to mark state as visited
			visitedTracker.SetVisited(currentState)
			ret := dfs(newState+i, depth+1)

			// Use VisitedTracker method to unmark state (backtracking)
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
	win := dfs(0, 0)
	fmt.Println("win:", win)
}

// game function allows user to participate in the game, with user input for even steps and algorithm choosing optimal solution for odd steps
func game() {
	// Create VisitedTracker instance
	visitedTracker := &VisitedTracker{0, 0}
	// Create memoization table
	dp := make(map[DPMemoKey]bool)
	state := 0
	depth := 0

	fmt.Println("Game started! You are the first player (even steps), please enter 0 or 1. The algorithm is the second player (odd steps).")
	fmt.Printf("Current state: %07b\n", state)

	for {
		var choice int
		if depth%2 == 0 {
			// Even step, user input
			fmt.Print("Please enter 0 or 1: ")
			fmt.Scan(&choice)
			if choice != 0 && choice != 1 {
				fmt.Println("Invalid input, please re-enter 0 or 1.")
				continue
			}
		} else {
			// Odd step, algorithm chooses optimal solution
			choice = findBestMove(state, depth, visitedTracker, dp)
			fmt.Printf("Algorithm chooses: %d\n", choice)
		}

		// Calculate new state
		newState := (state<<1)&(1<<7-1) + choice

		// Check for repetition
		if visitedTracker.IsVisited(newState) {
			if depth%2 == 0 {
				fmt.Println("Congratulations! You win!")
			} else {
				fmt.Println("Unfortunately, the algorithm wins!")
			}
			return
		}

		// Mark state as visited
		visitedTracker.SetVisited(newState)
		state = newState
		depth++

		fmt.Printf("Current state: %07b\n", state)
	}
}

// findBestMove helps the algorithm find the optimal solution
func findBestMove(state, depth int, visitedTracker *VisitedTracker, dp map[DPMemoKey]bool) int {
	// Create a temporary copy of visited states to avoid modifying the original
	tempVisited := &VisitedTracker{
		visited0: visitedTracker.visited0,
		visited1: visitedTracker.visited1,
	}

	var dfs func(int, int) (DFS_RESULT, int)
	dfs = func(s, d int) (DFS_RESULT, int) {
		v0, v1 := tempVisited.GetVisitedStates()
		key := DPMemoKey{
			state:    uint16(s),
			visited0: v0,
			visited1: v1,
		}

		if win, ok := dp[key]; ok {
			if win {
				return DFS_RESULT(d % 2), -1
			} else {
				return DFS_RESULT(1 - d%2), -1
			}
		}

		newState := (s << 1) & (1<<7 - 1)

		for i := 0; i < 2; i++ {
			currentState := newState + i

			if tempVisited.IsVisited(currentState) {
				return DFS_RESULT(d % 2), i
			}

			tempVisited.SetVisited(currentState)
			result, _ := dfs(newState+i, d+1)
			tempVisited.UnsetVisited(currentState)

			if result == DFS_RESULT(d%2) {
				return result, i
			}
		}

		result := DFS_RESULT(1 - d%2)
		return result, 0
	}

	_, bestMove := dfs(state, depth)
	return bestMove
}

func updateDPStatus(dp map[DPMemoKey]bool, key DPMemoKey, win bool) {
	fmt.Printf("updateDPStatus: visited: %064b%064b, state: %b, win: %v\n", key.visited1, key.visited0, key.state, win)
	dp[key] = win
}

func main() {
	fmt.Println("vim-go")
	// solution() // Commented out the original solution call
	game() // Call the new game function
}

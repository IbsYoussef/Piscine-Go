package solutions

import "github.com/01-edu/z01"

// Function Breakdown:
//
// EightQueens():
//   - Main function that initiates the solving process
//   - Prints all 92 solutions to the eight queens puzzle
//   - Uses backtracking and recursion
//
// Approach:
//   - Create board representation (array/slice)
//   - Start recursive backtracking from column 0
//   - For each column, try each row
//   - Check if placement is safe
//   - If safe, place queen and recurse to next column
//   - If all queens placed, print solution
//   - Backtrack and try next position
//
// Helper Functions Needed:
//   - isSafe(board, row, col): Check if placement is valid
//   - solve(board, col): Recursive backtracking function
//   - printSolution(board): Print board as required format
//
// Key Concepts:
//   - Backtracking: Core algorithm
//   - Recursion: Column by column placement
//   - Constraint checking: No attacking queens
//   - State management: Board representation
//
// Algorithm:
//   1. Try placing queen in each row of current column
//   2. Check if safe (no conflicts)
//   3. If safe, place queen and recurse to next column
//   4. If all 8 queens placed, print solution
//   5. Remove queen (backtrack) and try next row
//   6. Repeat until all possibilities explored
//
// Validation:
//   - Check row conflicts
//   - Check diagonal conflicts (both directions)
//   - Columns handled by recursion (one queen per column)

// Status: Bonus

func EightQueens() {
	board := make([]int, 8)
	solve(board, 0)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isSafe(board []int, row, col int) bool {
	// Check all previously placed queens (columns 0 to col-1)
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := board[prevCol]

		// Check same row
		if prevRow == row {
			return false
		}

		// Check Diagonals
		if abs(prevRow-row) == abs(prevCol-col) {
			return false
		}
	}

	return true
}

func printSolution(board []int) {
	for _, row := range board {
		z01.PrintRune(rune('0' + row + 1)) // Convert 0-7 to 1-8
	}
	z01.PrintRune('\n')
}

func solve(board []int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

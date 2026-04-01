package main

import (
	"fmt"
	"strconv"
)

// ============================================
// LeetCode 36: Valid Sudoku (Medium)
// ============================================
// https://leetcode.com/problems/valid-sudoku/
//
// Determine if a 9x9 Sudoku board is valid. Only the filled cells
// need to be validated according to the following rules:
// 1. Each row must contain the digits 1-9 without repetition. // Mỗi hàng duy nhất chứa số từ 1-9 và không lặp lại
// 2. Each column must contain the digits 1-9 without repetition. // mỗi cột duy nhất chứa số từ 1-9 và không lặp lại
// 3. Each of the nine 3x3 sub-boxes of the grid must contain the
//    digits 1-9 without repetition. // mỗi ô 3x3 duy nhất chứa số từ 1-9 và không lặp lại
//

func isValidSudoku(board [][]byte) bool {
	var rowSeen [9][10]bool
	var colSeen [9][10]bool
	var boxSeen [9][10]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num, _ := strconv.Atoi(string(board[i][j])) //'9' -> 9

			if rowSeen[i][num] {
				return false
			}
			rowSeen[i][num] = true

			if colSeen[j][num] {
				return false
			}
			colSeen[j][num] = true

			boxIdx := (i/3)*3 + (j / 3)
			if boxSeen[boxIdx][num] {
				return false
			}
			boxSeen[boxIdx][num] = true

		}
	}
	// TODO: Implement validation logic
	// Strategy: Use 3 sets of marker arrays (Rows, Cols, Boxes)

	return true
}

func main() {
	fmt.Println("=== LeetCode 36: Valid Sudoku ===")

	// Test Case 1: Valid board
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Board 1 is valid: %v (Expected: true)\n", isValidSudoku(board1))

	// Test Case 2: Invalid row (two 8s in the first column)
	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("Board 2 is valid: %v (Expected: false)\n", isValidSudoku(board2))
}

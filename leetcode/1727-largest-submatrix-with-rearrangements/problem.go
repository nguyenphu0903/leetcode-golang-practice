package main

import (
	"fmt"
	"sort"
)

// ============================================
// LeetCode 1727: Largest Submatrix With Rearrangements (Medium)
// ============================================
// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
//
// You are given a binary matrix matrix of size m x n. You are allowed to
// rearrange the columns of the matrix in any order.
//
// Return the area of the largest submatrix within matrix where every element
// of the submatrix is 1 after reordering the columns optimally.
//
// Example 1:
//   Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]
//   Output: 4
//   Explanation: You can rearrange the columns as shown above.
//   The largest submatrix of 1s, in bold, has an area of 4.
//
// Example 2:
//   Input: matrix = [[1,0,1,0,1]]
//   Output: 3
//   Explanation: You can rearrange the columns as shown above.
//   The largest submatrix of 1s, in bold, has an area of 3.
//
// Example 3:
//   Input: matrix = [[1,1,0],[1,0,1]]
//   Output: 2
//   Explanation: You can rearrange the columns as shown above.
//
// Constraints:
//   - m == matrix.length
//   - n == matrix[i].length
//   - 1 <= m * n <= 10^5
//   - matrix[i][j] is 0 or 1.

// TODO: Implement largestSubmatrix function
func largestSubmatrix(matrix [][]int) int {
	maxResult := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				continue
			}
			if matrix[i][j] == 1 {
				matrix[i][j] = matrix[i-1][j] + 1
			} else {
				matrix[i][j] = 0
			}
		}
		currRow := make([]int, len(matrix[i]))
		copy(currRow, matrix[i])
		sort.Sort(sort.Reverse(sort.IntSlice(currRow)))

		for k := 0; k < len(currRow); k++ {
			area := currRow[k] * (k + 1)
			if area > maxResult {
				maxResult = area
			}
		}
	}

	// Your code here
	// Time: O(M * N log N)
	// Space: O(N) or O(M * N)
	return maxResult
}

func main() {
	fmt.Println("=== LeetCode 1727: Largest Submatrix With Rearrangements ===\n")

	// Test Case 1
	matrix1 := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
	fmt.Printf("Input: matrix = %v\n", matrix1)
	fmt.Printf("Output: %d\n", largestSubmatrix(matrix1))
	fmt.Printf("Expected: 4\n\n")

	// Test Case 2
	matrix2 := [][]int{{1, 0, 1, 0, 1}}
	fmt.Printf("Input: matrix = %v\n", matrix2)
	fmt.Printf("Output: %d\n", largestSubmatrix(matrix2))
	fmt.Printf("Expected: 3\n\n")

	// Test Case 3
	matrix3 := [][]int{{1, 1, 0}, {1, 0, 1}}
	fmt.Printf("Input: matrix = %v\n", matrix3)
	fmt.Printf("Output: %d\n", largestSubmatrix(matrix3))
	fmt.Printf("Expected: 2\n\n")
}

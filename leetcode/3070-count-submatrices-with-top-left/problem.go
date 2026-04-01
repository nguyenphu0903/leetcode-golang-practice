package main

import "fmt"

// ============================================
// LeetCode 3070: Count Submatrices with Top-Left Element and Sum Less Than k (Medium)
// ============================================
// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
//
// You are given a 0-indexed m x n integer matrix grid and an integer k.
// Return the number of submatrices that contain the top-left element grid[0][0] 
// and have a sum less than or equal to k.
//
// Example 1:
//   Input: grid = [[7,6,3],[6,6,1]], k = 18
//   Output: 4
//   Explanation: There are 4 submatrices starting at (0,0) with sum <= 18:
//   - (0,0) to (0,0): sum = 7
//   - (0,0) to (0,1): sum = 7 + 6 = 13
//   - (0,0) to (0,2): sum = 7 + 6 + 3 = 16
//   - (0,0) to (1,0): sum = 7 + 6 = 13
//
// Example 2:
//   Input: grid = [[7,2,9],[1,5,0],[2,6,6]], k = 20
//   Output: 6
//
// Constraints:
//   - m == grid.length
//   - n == grid[i].length
//   - 1 <= m, n <= 1000
//   - 0 <= grid[i][j] <= 1000
//   - 1 <= k <= 10^9

// TODO: Implement countSubmatrices function
func countSubmatrices(grid [][]int, k int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Bước 1: Cộng dồn theo Hàng (Row-wise prefix sum)
	for i := 0; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += grid[i][j-1]
		}
	}

	// Bước 2: Cộng dồn theo Cột (Column-wise prefix sum) & Đếm
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
			// Bước 3: Đếm
			if grid[i][j] <= k {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println("=== LeetCode 3070: Count Submatrices with Top-Left Element and Sum Less Than k ===")

	// Test Case 1
	grid1 := [][]int{{7, 6, 3}, {6, 6, 1}}
	k1 := 18
	fmt.Printf("Input: grid = %v, k = %d\n", grid1, k1)
	fmt.Printf("Output: %d\n", countSubmatrices(grid1, k1))
	fmt.Printf("Expected: 4\n\n")

	// Test Case 2
	grid2 := [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}
	k2 := 20
	fmt.Printf("Input: grid = %v, k = %d\n", grid2, k2)
	fmt.Printf("Output: %d\n", countSubmatrices(grid2, k2))
	fmt.Printf("Expected: 6\n\n")
}

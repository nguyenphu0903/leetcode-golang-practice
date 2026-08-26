package main

import "fmt"

// ============================================
// LeetCode 167: Two Sum II - Input Array Is Sorted (Medium)
// ============================================
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
//
// Given a 1-indexed array of integers numbers that is already sorted
// in non-decreasing order, find two numbers such that they add up
// to a specific target number.
//
// Return the indices of the two numbers (1-indexed) as an integer
// array [index1, index2] of length 2, where 1 <= index1 < index2 <= len(numbers).
//
// The tests are generated such that there is exactly one solution.
// You may not use the same element twice.
//
// Your solution must use only constant extra space (O(1) memory).
//
// Example 1:
//   Input: numbers = [2,7,11,15], target = 9
//   Output: [1,2]
//   Explanation: 2 + 7 = 9 → indices 1 and 2 (1-indexed)
//
// Example 2:
//   Input: numbers = [2,3,4], target = 6
//   Output: [1,3]
//   Explanation: 2 + 4 = 6 → indices 1 and 3 (1-indexed)
//
// Example 3:
//   Input: numbers = [-1,0], target = -1
//   Output: [1,2]
//   Explanation: -1 + 0 = -1 → indices 1 and 2 (1-indexed)
//
// Constraints:
//   - 2 <= numbers.length <= 3 * 10^4
//   - -1000 <= numbers[i] <= 1000
//   - numbers is sorted in non-decreasing order
//   - -1000 <= target <= 1000
//   - Exactly one solution exists
//

func twoSumII(numbers []int, target int) []int {
	i := 0;
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1} // Return 1-indexed positions
		}
		if sum < target {
			i++
			continue
		}
		if sum > target {
			j--
			continue
		}
	}
	return []int{} // This line should never be reached due to problem constraints
}

func main() {
	fmt.Println("=== LeetCode 167: Two Sum II - Input Array Is Sorted ===")

	// Test Case 1
	nums1, target1 := []int{2, 7, 11, 15}, 9
	fmt.Printf("Input: numbers = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v (Expected: [1 2])\n\n", twoSumII(nums1, target1))

	// Test Case 2
	nums2, target2 := []int{2, 3, 4}, 6
	fmt.Printf("Input: numbers = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v (Expected: [1 3])\n\n", twoSumII(nums2, target2))

	// Edge Case 3 (negative numbers, smallest size)
	nums3, target3 := []int{-1, 0}, -1
	fmt.Printf("Input: numbers = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v (Expected: [1 2])\n", twoSumII(nums3, target3))
}

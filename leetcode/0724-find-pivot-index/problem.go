package main

import "fmt"

// ============================================
// LeetCode 724: Find Pivot Index (Easy)
// ============================================
// https://leetcode.com/problems/find-pivot-index/
//
// The pivot index is the index where the sum of all the numbers strictly to the
// left of the index is equal to the sum of all the numbers strictly to the
// index's right.
//
// If the index is on the left edge of the array, then the left sum is 0
// because there are no elements to the left. This also applies to the right
// edge of the array.
//
// Return the leftmost pivot index. If no such index exists, return -1.
//
// Example 1:
//   Input: nums = [1,7,3,6,5,6]
//   Output: 3
//   Explanation:
//   The pivot index is 3.
//   Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
//   Right sum = nums[4] + nums[5] = 5 + 6 = 11
//
// Example 2:
//   Input: nums = [1,2,3]
//   Output: -1
//
// Example 3:
//   Input: nums = [2,1,-1]
//   Output: 0
//   Explanation:
//   Left sum = 0, Right sum = nums[1] + nums[2] = 1 + -1 = 0
//
// Constraints:
//   - 1 <= nums.length <= 10^4
//   - -1000 <= nums[i] <= 1000
//

// TODO: Implement pivotIndex function
func pivotIndex(nums []int) int {
	arrRight := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			arrRight[i] = nums[i]
		} else {
			arrRight[i] = nums[i] + arrRight[i-1]
		}
	}
	arrLeft := make([]int, len(nums))
	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			arrLeft[j] = nums[j]
		} else {
			arrLeft[j] = nums[j] + arrLeft[j+1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if arrRight[i] == arrLeft[i] { // Chỉ so sánh tại đúng vị trí i
			return i
		}
	}
	// Your code here
	// Time: O(?)
	// Space: O(?)
	return -1
}

func main() {
	fmt.Println("=== LeetCode 724: Find Pivot Index ===")

	// Test Case 1
	nums1 := []int{1, 7, 3, 6, 5, 6}
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %d\n", pivotIndex(nums1))
	fmt.Printf("Expected: 3\n\n")

	// Test Case 2
	nums2 := []int{1, 2, 3}
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %d\n", pivotIndex(nums2))
	fmt.Printf("Expected: -1\n\n")

	// Test Case 3
	nums3 := []int{2, 1, -1}
	fmt.Printf("Input: nums = %v\n", nums3)
	fmt.Printf("Output: %d\n", pivotIndex(nums3))
	fmt.Printf("Expected: 0\n")
}

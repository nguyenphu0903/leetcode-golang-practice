package main

import "fmt"

// ============================================
// LeetCode 217: Contains Duplicate (Easy)
// ============================================
// https://leetcode.com/problems/contains-duplicate/
//
// Given an integer array nums, return true if any value appears
// at least twice in the array, and return false if every element is distinct.
//
// Example 1:
//   Input: nums = [1,2,3,1]
//   Output: true
//
// Example 2:
//   Input: nums = [1,2,3,4]
//   Output: false
//
// Example 3:
//   Input: nums = [1,1,1,3,3,4,3,2,4,2]
//   Output: true
//
// Constraints:
//   - 1 <= nums.length <= 10^5
//   - -10^9 <= nums[i] <= 10^9
//
// TODO: Implement function
func containsDuplicate(nums []int) bool {
	// Your code here
	// Hint: Use a map/set to track seen numbers
	seen := make(map[int]bool)
	for i := range nums {
		if seen[nums[i]] {
			return true
		}
		seen[nums[i]] = true

	}
	// Time: O(n) duyệt qua tất cả phần tử trong array
	// Space: O(n) do map lưu tất cả phần tử trong array
	return false
}

func main() {
	fmt.Println("=== LeetCode 217: Contains Duplicate ===\n")

	// Test case 1
	nums1 := []int{1, 2, 3, 1}
	result1 := containsDuplicate(nums1)
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Expected: true\n\n")

	// Test case 2
	nums2 := []int{1, 2, 3, 4}
	result2 := containsDuplicate(nums2)
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Expected: false\n\n")

	// Test case 3
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result3 := containsDuplicate(nums3)
	fmt.Printf("Input: nums = %v\n", nums3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Printf("Expected: true\n")
}


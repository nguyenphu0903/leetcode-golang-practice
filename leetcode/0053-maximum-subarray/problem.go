package main

import "fmt"

// ============================================
// LeetCode 53: Maximum Subarray (Easy)
// ============================================
// https://leetcode.com/problems/maximum-subarray/
//
// Given an integer array nums, find the contiguous subarray (containing
// at least one number) which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.
//
// Example 1:
//   Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//   Output: 6
//   Explanation: [4,-1,2,1] has the largest sum = 6.
//
// Example 2:
//   Input: nums = [1]
//   Output: 1
//
// Example 3:
//   Input: nums = [5,4,-1,7,8]
//   Output: 23
//
// Constraints:
//   - 1 <= nums.length <= 10^5
//   - -10^4 <= nums[i] <= 10^4
//
// Follow-up: If you have figured out the O(n) solution, try coding
//            another solution using the divide and conquer approach,
//            which is more subtle.
//
// TODO: Implement function
func maxSubArray(nums []int) int {
	// Your code here
	// Hint: Kadane's algorithm - track maximum sum ending at current position
	maxSum := nums[0]
    for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
	// Time: O(n^2) duyệt qua tất cả phần tử trong array
	// Space: O(1) chỉ dùng biến maxSum, không dùng thêm space
}

func maxSubArray2(nums []int) int {
	maxEndingHere := nums[0]
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxEndingHere < 0 {
			maxEndingHere = nums[i]
		} else {
			maxEndingHere += nums[i]
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
	// Time: O(n) duyệt qua tất cả phần tử trong array
	// Space: O(1) chỉ dùng biến maxEndingHere và maxSoFar, không dùng thêm space
}

func main() {
	fmt.Println("=== LeetCode 53: Maximum Subarray ===\n")

	// Test case 1
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result1 := maxSubArray(nums1)
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 6\n\n")

	// Test case 2
	nums2 := []int{1}
	result2 := maxSubArray(nums2)
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 1\n\n")

	// Test case 3
	nums3 := []int{5, 4, -1, 7, 8}
	result3 := maxSubArray(nums3)
	fmt.Printf("Input: nums = %v\n", nums3)
	fmt.Printf("Output: %d\n", result3)
	fmt.Printf("Expected: 23\n")
}


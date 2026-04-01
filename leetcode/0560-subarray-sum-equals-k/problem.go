package main

import "fmt"

// ============================================
// LeetCode 560: Subarray Sum Equals K (Medium)
// ============================================
// https://leetcode.com/problems/subarray-sum-equals-k/
//
// Given an array of integers nums and an integer k, return the total number
// of subarrays whose sum equals to k.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Example 1:
//   Input: nums = [1,1,1], k = 2
//   Output: 2
//
// Example 2:
//   Input: nums = [1,2,3], k = 3
//   Output: 2
//
// Constraints:
//   - 1 <= nums.length <= 2 * 10^4
//   - -1000 <= nums[i] <= 1000
//   - -10^7 <= k <= 10^7
//

// TODO: Implement subarraySum function
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	// Your code here
	// Hint: Use a Hash Map to store (prefixSum, count)
	// Time: O(?)
	// Space: O(?)
	return 0
}

func main() {
	fmt.Println("=== LeetCode 560: Subarray Sum Equals K ===")

	// Test Case 1
	nums1, k1 := []int{1, 1, 1}, 2
	fmt.Printf("Input: nums = %v, k = %d\n", nums1, k1)
	fmt.Printf("Output: %d\n", subarraySum(nums1, k1))
	fmt.Printf("Expected: 2\n\n")

	// Test Case 2
	nums2, k2 := []int{1, 2, 3}, 3
	fmt.Printf("Input: nums = %v, k = %d\n", nums2, k2)
	fmt.Printf("Output: %d\n", subarraySum(nums2, k2))
	fmt.Printf("Expected: 2\n\n")

	// Test Case 3 (Negative numbers)
	nums3, k3 := []int{1, -1, 0}, 0
	fmt.Printf("Input: nums = %v, k = %d\n", nums3, k3)
	fmt.Printf("Output: %d\n", subarraySum(nums3, k3))
	fmt.Printf("Expected: 3\n") // [1, -1], [0], [1, -1, 0]
}

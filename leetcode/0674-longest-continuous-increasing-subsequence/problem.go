package main

import (
	"fmt"
)

// ============================================
// LeetCode 674: Longest Continuous Increasing Subsequence (Easy)
// ============================================
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/
//
// Given an unsorted array of integers nums, return the length of the 
// longest continuous increasing subsequence.
// The subsequence must be strictly increasing and continuous (adjacent).
//
// Example 1:
//   Input: nums = [1,3,5,4,7]
//   Output: 3
//   Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
//                Even though [1,3,5,7] is increasing, it is not continuous.
//
// Example 2:
//   Input: nums = [2,2,2,2,2]
//   Output: 1
//   Explanation: [2] is the longest continuous increasing subsequence.
//
// Constraints:
//   - 1 <= nums.length <= 10^4
//   - -10^9 <= nums[i] <= 10^9
//

func findLengthOfLCIS(nums []int) int {
	// TODO: Implement logic
	// Gợi ý: Rất giống 485!
	//   - current: độ dài dãy tăng hiện tại
	currentLen := 1
	//   - longest: độ dài dãy tăng dài nhất
	longestLen := 1
	//   - longest: max
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLen++
			continue
		}
		if currentLen > longestLen {
			longestLen = currentLen
		}
		currentLen = 1
	}
	if currentLen > longestLen {
		longestLen = currentLen
	}
	//   - Duyệt từ i=1 đến hết:
	//       nếu nums[i] > nums[i-1] -> current++ (vẫn đang tăng)
	//       nếu không -> longest = max(longest, current), current = 1 (reset về 1 vì mỗi số tự nó là dãy dài 1)
	//   - Cuối: return max(longest, current)
	return longestLen
}

func main() {
	fmt.Println("=== LeetCode 674: Longest Continuous Increasing Subsequence ===")

	// Test Case 1
	nums1 := []int{1, 3, 5, 4, 7}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %d (Expected: 3)\n\n", findLengthOfLCIS(nums1))

	// Test Case 2
	nums2 := []int{2, 2, 2, 2, 2}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %d (Expected: 1)\n\n", findLengthOfLCIS(nums2))

	// Edge Case 3
	nums3 := []int{1, 3, 5, 7}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %d (Expected: 4)\n", findLengthOfLCIS(nums3))
}

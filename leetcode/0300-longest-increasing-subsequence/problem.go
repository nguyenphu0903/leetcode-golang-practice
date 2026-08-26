package main

import (
	"fmt"
)

// ============================================
// LeetCode 300: Longest Increasing Subsequence (Medium)
// ============================================
// https://leetcode.com/problems/longest-increasing-subsequence/
//
// Given an integer array nums, return the length of the longest
// strictly increasing subsequence.
//
// A subsequence is a sequence that can be derived from an array by
// deleting some or no elements without changing the order of the
// remaining elements. For example, [3,6,2,7] is a subsequence of [0,3,1,6,2,2,7].
//
// Example 1:
//   Input: nums = [10,9,2,5,3,7,101,18]
//   Output: 4
//   Explanation: The longest increasing subsequence is [2,3,7,101], therefore length is 4.
//
// Example 2:
//   Input: nums = [0,1,0,3,2,3]
//   Output: 4
//
// Example 3:
//   Input: nums = [7,7,7,7,7,7,7]
//   Output: 1
//
// Constraints:
//   - 1 <= nums.length <= 2500
//   - -10^4 <= nums[i] <= 10^4
//
// Follow up: Can you come up with O(n log n) solution?
//

func lengthOfLIS(nums []int) int {
	// TODO: Implement logic
	// Gợi ý: Khác 674 ở chỗ KHÔNG cần liên tục, được nhảy cóc!
	// Cách 1: DP O(n²) - dễ hiểu cho người mới:
	//   - dp[i] = độ dài LIS kết thúc tại i
	//   - dp[i] = 1 + max(dp[j]) với j < i và nums[j] < nums[i]
	//   - Khởi tạo dp[i] = 1 (mỗi số tự nó là dãy dài 1)
	//   - Đáp án = max(dp)
	//
	// Cách 2: Binary Search O(n log n) - nâng cao (follow up)
	//   - Dùng mảng tails, binary search vị trí thay thế
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	longest := 0
	for _, v := range dp {
		if v > longest {
			longest = v
		}
	}
	return longest
}

func main() {
	fmt.Println("=== LeetCode 300: Longest Increasing Subsequence ===")

	// Test Case 1
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %d (Expected: 4)\n\n", lengthOfLIS(nums1))

	// Test Case 2
	nums2 := []int{0, 1, 0, 3, 2, 3}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %d (Expected: 4)\n\n", lengthOfLIS(nums2))

	// Test Case 3
	nums3 := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %d (Expected: 1)\n", lengthOfLIS(nums3))
}

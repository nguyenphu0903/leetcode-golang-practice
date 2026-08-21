package main

import (
	"fmt"
)

// ============================================
// LeetCode 485: Max Consecutive Ones (Easy)
// ============================================
// https://leetcode.com/problems/max-consecutive-ones/
//
// Given a binary array nums, return the maximum number of consecutive 1's
// in the array.
//
// Example 1:
//   Input: nums = [1,1,0,1,1,1]
//   Output: 3
//   Explanation: The first two digits or the last three digits are consecutive 1s.
//                The maximum number of consecutive 1s is 3.
//
// Example 2:
//   Input: nums = [1,0,1,1,0,1]
//   Output: 2
//
// Constraints:
//   - 1 <= nums.length <= 10^5
//   - nums[i] is either 0 or 1.
//

func findMaxConsecutiveOnes(nums []int) int {
	// TODO: Implement logic
	// Gợi ý: Dùng 2 biến
	//   - current: đếm số 1 liên tiếp hiện tại
	//   - longest: lưu max
	// Duyệt từng số:
	//   nếu nums[i] == 1 -> current++
	//   nếu nums[i] == 0 -> so sánh và reset current = 0
	// Cuối cùng return max(longest, current) vì mảng có thể kết thúc bằng 1

	return 0
}

func main() {
	fmt.Println("=== LeetCode 485: Max Consecutive Ones ===")

	// Test Case 1
	nums1 := []int{1, 1, 0, 1, 1, 1}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %d (Expected: 3)\n\n", findMaxConsecutiveOnes(nums1))

	// Test Case 2
	nums2 := []int{1, 0, 1, 1, 0, 1}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %d (Expected: 2)\n\n", findMaxConsecutiveOnes(nums2))

	// Edge Case
	nums3 := []int{0, 0, 0}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %d (Expected: 0)\n", findMaxConsecutiveOnes(nums3))
}

package main

import (
	"fmt"
	"math"
)

// ============================================
// LeetCode 334: Increasing Triplet Subsequence (Medium)
// ============================================
// https://leetcode.com/problems/increasing-triplet-subsequence/
//
// Given an integer array nums, return true if there exists a triple
// of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
// If no such indices exists, return false.
//
// Example 1:
//   Input: nums = [1,2,3,4,5]
//   Output: true
//   Explanation: Any triplet where i < j < k is valid.
//
// Example 2:
//   Input: nums = [5,4,3,2,1]
//   Output: false
//   Explanation: No triplet exists.
//
// Example 3:
//   Input: nums = [2,1,5,0,4,6]
//   Output: true
//   Explanation: The triplet (3, 4, 5) is valid because nums[3]==0 < nums[4]==4 < nums[5]==6.
//
// Constraints:
//   - 1 <= nums.length <= 5*10^5
//   - -2^31 <= nums[i] <= 2^31-1
//

func increasingTriplet(nums []int) bool {
	// TODO: Implement logic
	// Gợi ý: Không cần DP như 300! Chỉ cần 2 biến:
	//   - first: số nhỏ nhất đã gặp
	//   - second: số nhỏ thứ 2 (lớn hơn first nhưng nhỏ nhất có thể)
	// Duyệt từng num:
	//   nếu num <= first -> first = num (tìm số nhỏ hơn)
	//   nếu num <= second -> second = num (tìm số vừa vừa)
	//   nếu num > second -> tìm được triplet! return true
	// Hết vòng -> return false

	return false
}

func main() {
	fmt.Println("=== LeetCode 334: Increasing Triplet Subsequence ===")

	// Test Case 1
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %v (Expected: true)\n\n", increasingTriplet(nums1))

	// Test Case 2
	nums2 := []int{5, 4, 3, 2, 1}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %v (Expected: false)\n\n", increasingTriplet(nums2))

	// Test Case 3
	nums3 := []int{2, 1, 5, 0, 4, 6}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %v (Expected: true)\n", increasingTriplet(nums3))

	// Dùng math.MaxInt để khởi tạo first, second cho dễ (gợi ý)
	_ = math.MaxInt
}

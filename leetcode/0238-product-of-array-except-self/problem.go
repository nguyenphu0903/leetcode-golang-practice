package main

import "fmt"

// ============================================
// LeetCode 238: Product of Array Except Self (Medium)
// ============================================
// https://leetcode.com/problems/product-of-array-except-self/
//
// Given an integer array nums, return an array answer such that answer[i]
// is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
//
//	Input: nums = [1,2,3,4]
//	Output: [24,12,8,6]
//
// Example 2:
//
//	Input: nums = [-1,1,0,-3,3]
//	Output: [0,0,9,0,0]
//
// Constraints:
//   - 2 <= nums.length <= 10^5
//   - -30 <= nums[i] <= 30
//   - The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// Follow-up: Can you solve the problem in O(1) extra space complexity?
//
//	(The output array does not count as extra space for space complexity analysis.)
//
// TODO: Implement function
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]

	}
	// Step 1: Calculate prefix products (going from left to right)
	// ans[i] should store product of all elements to the left of nums[i]

	// Step 2: Multiply with suffix products (going from right to left)
	// You can use a single variable to keep track of suffix product to achieve O(1) space optimization.

	rightProduct := 1
	for j := n - 1; j >= 0; j-- {
		ans[j] *= rightProduct
		rightProduct *= nums[j]
	}
	return ans
}

func main() {
	fmt.Println("=== LeetCode 238: Product of Array Except Self ===\n")

	// Test case 1
	nums1 := []int{1, 2, 3, 4}
	result1 := productExceptSelf(nums1)
	fmt.Printf("Input: nums = %v\n", nums1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Expected: [24 12 8 6]\n\n")

	// Test case 2
	nums2 := []int{-1, 1, 0, -3, 3}
	result2 := productExceptSelf(nums2)
	fmt.Printf("Input: nums = %v\n", nums2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Expected: [0 0 9 0 0]\n")
}

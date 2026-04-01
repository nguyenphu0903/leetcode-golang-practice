package main

import (
	"fmt"
	"sort"
)

// ============================================
// LeetCode 347: Top K Frequent Elements (Medium)
// ============================================
// https://leetcode.com/problems/top-k-frequent-elements/
//
// Given an integer array nums and an integer k, return the k most
// frequent elements. You may return the answer in any order.
//
// Example 1:
//   Input: nums = [1,1,1,2,2,3], k = 2
//   Output: [1,2]
//
// Example 2:
//   Input: nums = [1], k = 1
//   Output: [1]
//
// Constraints:
//   - 1 <= nums.length <= 10^5
//   - k is in the range [1, number of unique elements in the array].
//   - It is guaranteed that the answer is unique.
//

// TODO: Implement topKFrequent function
func topKFrequent(nums []int, k int) []int {
	type Pair struct {
		num  int
		freq int
	}

	// Step 1: Count frequency using a Map (O(N))
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	// Step 2: Transfer unique numbers to a slice for sorting (O(M))
	uniquePairs := make([]Pair, 0, len(counts))
	for num, freq := range counts {
		uniquePairs = append(uniquePairs, Pair{num, freq})
	}

	// Step 3: Sort by frequency descending (O(M log M))
	sort.Slice(uniquePairs, func(i, j int) bool {
		return uniquePairs[i].freq > uniquePairs[j].freq
	})

	// Step 4: Collect top K
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, uniquePairs[i].num)
	}

	return result
}

func main() {
	fmt.Println("=== LeetCode 347: Top K Frequent Elements ===")

	// Test Case 1
	nums1, k1 := []int{1, 1, 1, 2, 2, 3}, 2
	fmt.Printf("Input: nums=%v, k=%d\n", nums1, k1)
	fmt.Printf("Output: %v\n", topKFrequent(nums1, k1))
	fmt.Println("Expected: [1 2]\n")

	// Test Case 2
	nums2, k2 := []int{1}, 1
	fmt.Printf("Input: nums=%v, k=%d\n", nums2, k2)
	fmt.Printf("Output: %v\n", topKFrequent(nums2, k2))
	fmt.Println("Expected: [1]\n")
}

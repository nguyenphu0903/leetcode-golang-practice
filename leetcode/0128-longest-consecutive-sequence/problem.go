package main

import (
	"fmt"
)

// ============================================
// LeetCode 128: Longest Consecutive Sequence (Medium)
// ============================================
// https://leetcode.com/problems/longest-consecutive-sequence/
//
// Given an unsorted array of integers nums, return the length of 
// the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
// Example 1:
//   Input: nums = [100,4,200,1,3,2]
//   Output: 4
//
// Example 2:
//   Input: nums = [0,3,7,2,5,8,4,6,0,1]
//   Output: 9
//

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Strategy: 1. Put all nums into a Map (Set)
	mapSet := make(map[int]bool)
	for _, num := range nums {
		mapSet[num] = true
	}

	// Strategy: 2 & 3. Find start of each sequence and count
	longestStreak := 0
	for num := range mapSet {
		// Only count if num is the start of a sequence
		if _, exists := mapSet[num-1]; exists {
			continue
		}
		// num is the start of a sequence
		currentNum := num
		currentStreak := 1
		for mapSet[currentNum+1] {
			currentNum++
			currentStreak++
		}
		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}

	return longestStreak
}

func main() {
	fmt.Println("=== LeetCode 128: Longest Consecutive Sequence ===")

	// Test Case 1
	nums1 := []int{100, 4, 200, 1, 3, 2}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %d (Expected: 4)\n\n", longestConsecutive(nums1))

	// Test Case 2
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %d (Expected: 9)\n", longestConsecutive(nums2))
}

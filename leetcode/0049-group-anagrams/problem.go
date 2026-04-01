package main

import (
	"fmt"
)

// ============================================
// LeetCode 49: Group Anagrams (Medium)
// ============================================
// https://leetcode.com/problems/group-anagrams/
//
// Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters
// of a different word or phrase, typically using all the original
// letters exactly once.
//
// Example 1:
//   Input: strs = ["eat","tea","tan","ate","nat","bat"]
//   Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
//   Input: strs = [""]
//   Output: [[""]]
//
// Example 3:
//   Input: strs = ["a"]
//   Output: [["a"]]
//
// Constraints:
//   - 1 <= strs.length <= 10^4
//   - 0 <= strs[i].length <= 100
//   - strs[i] consists of lowercase English letters.
//

// TODO: Implement groupAnagrams function
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		count := [26]int{}
		for _, char := range strs[i] {
			count[char-'a']++
		}
		groups[count] = append(groups[count], strs[i])
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	fmt.Println("=== LeetCode 49: Group Anagrams ===")

	// Test Case 1
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Output: %v\n", groupAnagrams(strs1))
	fmt.Printf("Expected: [[bat] [nat tan] [ate eat tea]] (order may vary)\n\n")

	// Test Case 2
	strs2 := []string{""}
	fmt.Printf("Input: %v\n", strs2)
	fmt.Printf("Output: %v\n", groupAnagrams(strs2))
	fmt.Printf("Expected: [[]]\n\n")

	// Test Case 3
	strs3 := []string{"a"}
	fmt.Printf("Input: %v\n", strs3)
	fmt.Printf("Output: %v\n", groupAnagrams(strs3))
	fmt.Printf("Expected: [[a]]\n")
}

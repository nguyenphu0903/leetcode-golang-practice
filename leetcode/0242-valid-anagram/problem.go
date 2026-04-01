package main

import "fmt"

// ============================================
// LeetCode 242: Valid Anagram (Easy)
// ============================================
// https://leetcode.com/problems/valid-anagram/
//
// Given two strings s and t, return true if t is an anagram of s,
// and false otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters
// of a different word or phrase, typically using all the original
// letters exactly once.
//
// Example 1:
//
//	Input: s = "anagram", t = "nagaram"
//	Output: true
//
// Example 2:
//
//	Input: s = "rat", t = "car"
//	Output: false
//
// Constraints:
//   - 1 <= s.length, t.length <= 5 * 10^4
//   - s and t consist of lowercase English letters.
//
// Follow-up: What if the inputs contain Unicode characters?
//
//	How would you adapt your solution to such a case?
//
// TODO: Implement function
func isAnagram(s string, t string) bool {
	// Your code here
	// Hint: Count character frequencies in both strings
	if len(s) != len(t) {
		return false
	}
	countS := make(map[rune]int)
	countT := make(map[rune]int)
	for _, v := range s {
		countS[v]++
	}
	for _, v := range t {
		countT[v]++
	}
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
	// Time: O(n) duyệt qua tất cả phần tử trong string
	// Space: O(n) do map lưu tất cả phần tử trong string
}

func main() {
	fmt.Println("=== LeetCode 242: Valid Anagram ===\n")

	// Test case 1
	s1, t1 := "anagram", "nagaram"
	result1 := isAnagram(s1, t1)
	fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", s1, t1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Expected: true\n\n")

	// Test case 2
	s2, t2 := "rat", "car"
	result2 := isAnagram(s2, t2)
	fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", s2, t2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Expected: false\n\n")

	// Test case 3
	s3, t3 := "listen", "silent"
	result3 := isAnagram(s3, t3)
	fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", s3, t3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Printf("Expected: true\n\n")

	// Test case 4 (Regression: Subset check)
	s4, t4 := "a", "ab"
	result4 := isAnagram(s4, t4)
	fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", s4, t4)
	fmt.Printf("Output: %v\n", result4)
	fmt.Printf("Expected: false\n")
}

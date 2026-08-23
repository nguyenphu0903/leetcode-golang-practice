package main

import "fmt"

// ============================================
// LeetCode 383: Ransom Note (Easy)
// ============================================
// https://leetcode.com/problems/ransom-note/
//
// Given two strings ransomNote and magazine, return true if ransomNote
// can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
//
// Example 1:
//   Input: ransomNote = "a", magazine = "b"
//   Output: false
//
// Example 2:
//   Input: ransomNote = "aa", magazine = "ab"
//   Output: false
//
// Example 3:
//   Input: ransomNote = "aa", magazine = "aab"
//   Output: true
//
// Constraints:
//   - 1 <= ransomNote.length, magazine.length <= 10^5
//   - Both strings consist of lowercase English letters
//

func canConstruct(ransomNote string, magazine string) bool {
	// Gợi ý: Cùng dạng "đếm tần suất ký tự" như Valid Anagram (242)!
	// Chọn Cách 2 (mảng cố định) vì đề cho biết chỉ có chữ cái thường a-z.
	//   - freq := [26]int{}
	//   - index = ký tự - 'a'
	// Time: O(n + m) với n, m là độ dài 2 chuỗi
	// Space: O(1) — đúng 26 ô cố định
	freq := [26]int{}
	for _, kitu := range magazine {
		freq[kitu-'a']++
	}
	for _, kitu := range ransomNote {
		if freq[kitu-'a'] == 0 {
			return false
		}
		freq[kitu-'a']--
	}
	return true
}

func main() {
	fmt.Println("=== LeetCode 383: Ransom Note ===")

	// Test Case 1
	r1, m1 := "a", "b"
	fmt.Printf("Input: ransomNote = \"%s\", magazine = \"%s\"\n", r1, m1)
	fmt.Printf("Output: %v (Expected: false)\n\n", canConstruct(r1, m1))

	// Test Case 2
	r2, m2 := "aa", "ab"
	fmt.Printf("Input: ransomNote = \"%s\", magazine = \"%s\"\n", r2, m2)
	fmt.Printf("Output: %v (Expected: false)\n\n", canConstruct(r2, m2))

	// Test Case 3
	r3, m3 := "aa", "aab"
	fmt.Printf("Input: ransomNote = \"%s\", magazine = \"%s\"\n", r3, m3)
	fmt.Printf("Output: %v (Expected: true)\n\n", canConstruct(r3, m3))

	// Edge Case 4 (same single char)
	r4, m4 := "a", "a"
	fmt.Printf("Input: ransomNote = \"%s\", magazine = \"%s\"\n", r4, m4)
	fmt.Printf("Output: %v (Expected: true)\n", canConstruct(r4, m4))
}

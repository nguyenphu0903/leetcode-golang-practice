package main

import "fmt"

// ============================================
// LeetCode 125: Valid Palindrome (Easy)
// ============================================
// https://leetcode.com/problems/valid-palindrome/
//
// A phrase is a palindrome if, after converting all uppercase letters
// into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
//   Input: s = "A man, a plan, a canal: Panama"
//   Output: true
//   Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
//   Input: s = "race a car"
//   Output: false
//   Explanation: "raceacar" is not a palindrome.
//
// Example 3:
//   Input: s = " "
//   Output: true
//   Explanation: After removing non-alphanumeric, s becomes "" (empty).
//                An empty string reads the same forward and backward.
//
// Constraints:
//   - 1 <= s.length <= 2 * 10^5
//   - s consists only of printable ASCII characters
//

func isPalindrome(s string) bool {
	// TODO: Implement logic
	// Gợi ý: Đây là bài TWO POINTERS đầu tiên của lộ trình!
	//
	//   Bước 1 - Làm sạch chuỗi:
	//     - Chỉ giữ chữ cái (a-z) và số (0-9), chữ hoa chuyển về chữ thường
	//     - Duyệt rune: nếu ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') -> giữ
	//     - Dùng unicode.ToLower(c) hoặc tự trừ 32 cho chữ HOA
	//
	//   Bước 2 - Hai con trỏ tiến về nhau:
	//     - left := 0, right := len(clean)-1
	//     - while left < right:
	//         + clean[left] != clean[right] ? -> return false
	//         + left++, right--
	//     - return true
	//
	// Time: O(n)
	// Space: O(n) cho chuỗi đã làm sạch
	cleanStr := make([]rune, 0, len(s))
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			cleanStr = append(cleanStr, c)
		} else if 'A' <= c && c <= 'Z' {
			cleanStr = append(cleanStr, c+'a'-'A') // Chuyển chữ hoa về chữ thường
		}
	}

	left, right := 0, len(cleanStr)-1
	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println("=== LeetCode 125: Valid Palindrome ===")

	// Test Case 1
	s1 := "A man, a plan, a canal: Panama"
	fmt.Printf("Input: %q\n", s1)
	fmt.Printf("Output: %v (Expected: true)\n\n", isPalindrome(s1))

	// Test Case 2
	s2 := "race a car"
	fmt.Printf("Input: %q\n", s2)
	fmt.Printf("Output: %v (Expected: false)\n\n", isPalindrome(s2))

	// Edge Case 3 (only spaces/punctuation)
	s3 := " "
	fmt.Printf("Input: %q\n", s3)
	fmt.Printf("Output: %v (Expected: true)\n\n", isPalindrome(s3))

	// Tricky Case 4 ('0' vs 'P' - digit vs letter)
	s4 := "0P"
	fmt.Printf("Input: %q\n", s4)
	fmt.Printf("Output: %v (Expected: false)\n", isPalindrome(s4))
}

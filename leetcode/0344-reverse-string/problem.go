package main

import "fmt"

// ============================================
// LeetCode 344: Reverse String (Easy)
// ============================================
// https://leetcode.com/problems/reverse-string/
//
// Write a function that reverses a string. The input string is given
// as an array of characters s.
//
// You must do this by modifying the input array in-place with O(1)
// extra memory.
//
// Example 1:
//   Input: s = ["h","e","l","l","o"]
//   Output: ["o","l","l","e","h"]
//
// Example 2:
//   Input: s = ["H","a","n","n","a","h"]
//   Output: ["h","a","n","n","a","H"]
//
// Constraints:
//   - 1 <= s.length <= 10^5
//

func reverseString(s []byte) {
	// TODO: Implement logic — LẦN NÀY BẠN TỰ CODE TOÀN BỘ!
	//
	// Gợi ý: Đúng khung TWO POINTERS hội tụ như Two Sum II vừa học!
	//   - left := 0, right := len(s)-1
	//   - while left < right:
	//       hoán đổi s[left] và s[right]
	//       left++, right--
	//
	// Cú pháp swap trong Go (không cần biến tạm):
	//   s[left], s[right] = s[right], s[left]
	//
	// Time: O(n)      — duyệt nửa mảng
	// Space: O(1)     — in-place, đúng yêu cầu đề bài
}

func printBytes(label string, b []byte) {
	fmt.Printf("%s: [", label)
	for i, c := range b {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%q", c)
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("=== LeetCode 344: Reverse String ===")

	// Test Case 1
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	printBytes("Input ", s1)
	reverseString(s1)
	printBytes("Output", s1)
	fmt.Println("Expected: ['o', 'l', 'l', 'e', 'h']\n")

	// Test Case 2 (palindrome-ish, even length)
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	printBytes("Input ", s2)
	reverseString(s2)
	printBytes("Output", s2)
	fmt.Println("Expected: ['h', 'a', 'n', 'n', 'a', 'H']\n")

	// Edge Case 3 (single char — không có gì để swap!)
	s3 := []byte{'a'}
	printBytes("Input ", s3)
	reverseString(s3)
	printBytes("Output", s3)
	fmt.Println("Expected: ['a']")
}

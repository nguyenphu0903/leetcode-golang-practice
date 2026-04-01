package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ============================================
// LeetCode 271: Encode and Decode Strings (Medium)
// ============================================
// https://leetcode.com/problems/encode-and-decode-strings/
//
// Example 1:
//   Input: ["lint","code","love","you"]
//   Output: ["lint","code","love","you"]
//
// Example 2:
//   Input: ["we", "say", ":", "yes"]
//   Output: ["we", "say", ":", "yes"]
//

// Encode encodes a list of strings to a single string.
func Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		// Format: [Length][#][Content]
		// Example: "lint" -> "4#lint"
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

// Decode decodes a single string to a list of strings.
func Decode(s string) []string {
	res := []string{}
	i := 0
	for i < len(s) {
		// Find the delimiter '#' to get the length
		j := i
		for s[j] != '#' {
			j++
		}
		
		// Parse the length
		length, _ := strconv.Atoi(s[i:j])
		
		// The string starts after '#'
		start := j + 1
		// It ends after 'length' characters
		end := start + length
		
		res = append(res, s[start:end])
		
		// Move pointer for the next word
		i = end
	}
	return res
}

func main() {
	fmt.Println("=== Encode and Decode Strings ===")

	// Test Case 1
	strs1 := []string{"lint", "code", "love", "you"}
	encoded1 := Encode(strs1)
	decoded1 := Decode(encoded1)
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Encoded: %s\n", encoded1)
	fmt.Printf("Decoded: %v\n", decoded1)
	fmt.Println("Result:", fmt.Sprint(strs1) == fmt.Sprint(decoded1))
	fmt.Println()

	// Test Case 2
	strs2 := []string{"we", "say", ":", "yes", "###", "4#hello"}
	encoded2 := Encode(strs2)
	decoded2 := Decode(encoded2)
	fmt.Printf("Input: %v\n", strs2)
	fmt.Printf("Encoded: %s\n", encoded2)
	fmt.Printf("Decoded: %v\n", decoded2)
	fmt.Println("Result:", fmt.Sprint(strs2) == fmt.Sprint(decoded2))
}

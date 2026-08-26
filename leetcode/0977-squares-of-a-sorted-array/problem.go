package main

import "fmt"

// ============================================
// LeetCode 977: Squares of a Sorted Array (Easy)
// ============================================
// https://leetcode.com/problems/squares-of-a-sorted-array/
//
// Given an integer array nums sorted in non-decreasing order, return
// an array of the squares of each number sorted in non-decreasing order.
//
// Follow up: squaring each element and sorting the new array is very
// easy (O(n log n)). Can you find an O(n) solution?
//
// Example 1:
//
//	Input: nums = [-4,-1,0,3,10]
//	Output: [0,1,9,16,100]
//	Explanation: squared = [16,1,0,9,100], after sorting -> [0,1,9,16,100]
//
// Example 2:
//
//	Input: nums = [-7,-3,2,3,11]
//	Output: [4,9,9,49,121]
//
// Constraints:
//   - 1 <= nums.length <= 10^4
//   - -10^4 <= nums[i] <= 10^4
//   - nums is sorted in non-decreasing order
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	// TODO: Implement logic
	//
	// Gợi ý 1 - Quan sát quan trọng:
	//   Mảng đã sort, nhưng có SỐ ÂM! Bình phương số âm thành số dương LỚN.
	//   => Phần tử bình phương LỚN NHẤT luôn nằm ở HAI ĐẦU mảng,
	//      KHÔNG bao giờ ở giữa!
	//   VD: [-4,-1,0,3,10] -> đầu trái ^2=16, đầu phải ^2=100, giữa nhỏ
	//
	// Gợi ý 2 - TWO POINTERS... nhưng lần này ĐẨY NGƯỢC từ cuối:
	//   - left := 0, right := len(nums)-1
	//   - Tạo mảng kết quả res cùng kích thước
	//   - Duy trì vị trí điền: i chạy từ CUỐI res về 0
	//   - Mỗi bước: so sánh |nums[left]| vs |nums[right]|
	//       bên nào BÌNH PHƯƠNG LỚN HƠN -> đặt vào res[i]
	//       rồi dịch con trỏ phía đó vào trong, i--
	//
	// Dry run thử: [-4,-1,0,3,10]
	//   |10| > |-4| -> res[4]=100, right--
	//   |-4| > |3|  -> res[3]=16,  left++
	//   ...
	//
	// Time: O(n) — mỗi phần tử được thăm đúng 1 lần
	// Space: O(n) cho mảng kết quả (đề bắt buộc phải trả về mảng mới)
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			res[right-left] = nums[left] * nums[left]
			left++
			continue
		}
		if abs(nums[left]) <= abs(nums[right]) {
			res[right-left] = nums[right] * nums[right]
			right--
			continue
		}
	}
	return res
}

func main() {
	fmt.Println("=== LeetCode 977: Squares of a Sorted Array ===")

	// Test Case 1 (negative + positive)
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %v\n", sortedSquares(nums1))
	fmt.Println("Expected: [0 1 9 16 100]\n")

	// Test Case 2
	nums2 := []int{-7, -3, 2, 3, 11}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %v\n", sortedSquares(nums2))
	fmt.Println("Expected: [4 9 9 49 121]\n")

	// Edge Case 3 (all negative)
	nums3 := []int{-5, -3}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %v\n", sortedSquares(nums3))
	fmt.Println("Expected: [9 25]")
}

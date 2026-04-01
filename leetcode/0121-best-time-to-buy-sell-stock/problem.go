package main

import "fmt"

// ============================================
// LeetCode 121: Best Time to Buy and Sell Stock (Easy)
// ============================================
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
// You are given an array prices where prices[i] is the price of a given stock
// on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.
//
// Example 1:
//   Input: prices = [7,1,5,3,6,4]
//   Output: 5
//   Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6),
//                profit = 6-1 = 5.
//                Note that buying on day 2 and selling on day 1 is not allowed
//                because you must buy before you sell.
//
// Example 2:
//   Input: prices = [7,6,4,3,1]
//   Output: 0
//   Explanation: In this case, no transactions are done and the max profit = 0.
//
// Constraints:
//   - 1 <= prices.length <= 10^5
//   - 0 <= prices[i] <= 10^4
//
// TODO: Implement function
func maxProfit(prices []int) int {
	// Your code here
	// Hint: Track minimum price seen so far, calculate profit for each day
	minPrice := prices[0]
	maxProfit := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i] - minPrice > maxProfit{
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
	// Time: O(n) duyệt qua tất cả phần tử trong array
	// Space: O(1) chỉ dùng biến minPrice và maxProfit, không dùng thêm space
}

func main() {
	fmt.Println("=== LeetCode 121: Best Time to Buy and Sell Stock ===\n")

	// Test case 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	result1 := maxProfit(prices1)
	fmt.Printf("Input: prices = %v\n", prices1)
	fmt.Printf("Output: %d\n", result1)
	fmt.Printf("Expected: 5\n\n")

	// Test case 2
	prices2 := []int{7, 6, 4, 3, 1}
	result2 := maxProfit(prices2)
	fmt.Printf("Input: prices = %v\n", prices2)
	fmt.Printf("Output: %d\n", result2)
	fmt.Printf("Expected: 0\n\n")

	// Test case 3
	prices3 := []int{2, 4, 1}
	result3 := maxProfit(prices3)
	fmt.Printf("Input: prices = %v\n", prices3)
	fmt.Printf("Output: %d\n", result3)
	fmt.Printf("Expected: 2\n")
}


// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 买卖股票的最佳时机 IV
package main

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if k >= (n >> 1) {
		return infiniteTran(prices)
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 1; j <= k; j++ {
		prev := -prices[0]
		for i := 1; i < n; i++ {
			dp[j][i] = max(dp[j][i-1], prev+prices[i])
			prev = max(prev, dp[j-1][i]-prices[i])
		}
	}
	return dp[k][n-1]
}

func infiniteTran(prices []int) (profit int) {
	for i := 1; i < len(prices); i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}
	return
}

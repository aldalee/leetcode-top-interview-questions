// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 买卖股票的最佳时机含冷冻期
package main

func maxProfit5(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	hold := make([]int, n) // 第[i]天结束时，持有股票时的最大收益
	cash := make([]int, n) // 第[i]天结束时，持有现金时的最大收益
	hold[1] = max(-prices[0], -prices[1])
	cash[1] = max(0, prices[1]-prices[0])
	for i := 2; i < n; i++ {
		hold[i] = max(hold[i-1], cash[i-2]-prices[i])
		cash[i] = max(cash[i-1], hold[i-1]+prices[i])
	}
	return cash[n-1]
}

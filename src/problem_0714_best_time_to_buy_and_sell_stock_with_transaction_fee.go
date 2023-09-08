// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
// 买卖股票的最佳时机含手续费
package main

func maxProfit6(prices []int, fee int) int {
	n := len(prices)
	hold := make([]int, n) // 第[i]天结束时，持有股票时的最大收益
	cash := make([]int, n) // 第[i]天结束时，持有现金时的最大收益
	hold[0] = -prices[0]
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], cash[i-1]-prices[i])
		cash[i] = max(cash[i-1], hold[i-1]+prices[i]-fee)
	}
	return cash[n-1]
}

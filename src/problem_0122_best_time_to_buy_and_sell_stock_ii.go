// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 买卖股票的最佳时机 II
package main

func maxProfit2(prices []int) (money int) {
	for i := 1; i < len(prices); i++ {
		money += max(prices[i]-prices[i-1], 0)
	}
	return
}

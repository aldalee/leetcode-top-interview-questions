// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 买卖股票的最佳时机
package main

func maxProfit(prices []int) (money int) {
	buy := prices[0]
	for _, sell := range prices {
		buy = min(buy, sell)
		money = max(money, sell-buy)
	}
	return
}

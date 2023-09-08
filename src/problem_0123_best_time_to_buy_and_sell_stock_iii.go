// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 买卖股票的最佳时机 III
package main

func maxProfit3(prices []int) (money int) {
	firstBuy := prices[0]              // 最佳买入价格
	onceProfMax := 0                   // 单笔交易最大利润
	onceProfSecondBuyMax := -prices[0] // 第二次买入后最大利润
	for i := 1; i < len(prices); i++ {
		firstBuy = min(firstBuy, prices[i])
		onceProfMax = max(onceProfMax, prices[i]-firstBuy)
		onceProfSecondBuyMax = max(onceProfSecondBuyMax, onceProfMax-prices[i])
		money = max(money, onceProfSecondBuyMax+prices[i])
	}
	return
}

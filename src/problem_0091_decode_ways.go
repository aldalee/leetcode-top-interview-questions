// https://leetcode.cn/problems/decode-ways/
// 解码方法
package main

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] = dp[i+1]
			if i+1 < n && (s[i]-'0')*10+(s[i+1]-'0') < 27 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

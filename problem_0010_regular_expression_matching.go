// https://leetcode.cn/problems/regular-expression-matching/
// 正则表达式匹配
package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				if matches(s, p, i, j) {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if matches(s, p, i, j-1) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}

func matches(s string, p string, i int, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}

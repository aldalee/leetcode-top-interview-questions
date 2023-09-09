// https://leetcode.cn/problems/palindrome-partitioning/
// 分割回文串
package main

func partition(s string) (res [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}
	var dfs func(idx int)
	var path []string
	dfs = func(idx int) {
		if idx == n {
			res = append(res, append(path[:0:0], path...))
			return
		}
		for step := idx; step < n; step++ {
			if dp[idx][step] {
				path = append(path, s[idx:step+1])
				dfs(step + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}

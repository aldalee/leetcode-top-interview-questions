// https://leetcode.cn/problems/word-break-ii
// 单词拆分 II
package main

import "strings"

func wordBreak2(s string, wordDict []string) []string {
	n := len(s)
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			if dp[j] && dict[s[i:j]] {
				dp[i] = true
				break
			}
		}
	}
	var (
		dfs  func(idx int)
		res  []string
		path []string
	)
	dfs = func(start int) {
		if start == n {
			res = append(res, strings.Join(path, " "))
			return
		}
		for end := start + 1; end <= n; end++ {
			prefix := s[start:end]
			if dict[prefix] && dp[end] {
				path = append(path, prefix)
				dfs(end)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

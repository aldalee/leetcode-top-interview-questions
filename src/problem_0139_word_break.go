// https://leetcode.cn/problems/word-break/
// 单词拆分
package main

import "github.com/emirpasic/gods/sets/hashset"

func wordBreak(s string, wordDict []string) bool {
	dict := hashset.New()
	for _, word := range wordDict {
		dict.Add(word)
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			if dp[j] && dict.Contains(s[i:j]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 无重复字符的最长子串
package main

func lengthOfLongestSubstring(s string) int {
	var m [256]int
	for i := 0; i < len(m); i++ {
		m[i] = -1
	}
	res := 0
	pre := -1
	for idx := 0; idx < len(s); idx++ {
		pre = max(pre, m[s[idx]])
		res = max(res, idx-pre)
		m[s[idx]] = idx
	}
	return res
}

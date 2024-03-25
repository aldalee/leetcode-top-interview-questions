// https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
// 至少有K个重复字符的最长子串
package main

import "strings"

func longestSubstring(s string, k int) (ans int) {
	var cnt = make([]int, 256)
	for _, ch := range s {
		cnt[ch]++
	}
	var sep byte
	for ch, n := range cnt {
		if n > 0 && n < k {
			sep = byte(ch)
			break
		}
	}
	if sep == 0 {
		return len(s)
	}
	for _, subStr := range strings.Split(s, string(sep)) {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return
}

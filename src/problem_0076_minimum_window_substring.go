// https://leetcode.cn/problems/minimum-window-substring
// 最小覆盖子串
package main

func minWindow(s, t string) string {
	m := [256]int{} // idx: key 'letter', m[idx]: value 'cnt'
	for _, c := range t {
		m[c]++
	}
	match, minLen := len(t), -1
	l, r := 0, 0   // the window is [l,r)
	L, R := -1, -1 // the answer is [L,R)
	for r < len(s) {
		m[s[r]]--
		if m[s[r]] >= 0 {
			match--
		}
		for match == 0 {
			if minLen == -1 || minLen > r-l+1 {
				minLen = r - l + 1
				L, R = l, r
			}
			m[s[l]]++
			if m[s[l]] > 0 {
				match++
			}
			l++
		}
		r++
	}
	if minLen != -1 {
		return s[L : R+1]
	}
	return ""
}

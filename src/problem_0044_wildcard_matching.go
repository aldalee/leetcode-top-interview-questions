// https://leetcode.cn/problems/wildcard-matching
// 通配符匹配
package main

// IsMatch please modify the function name to "isMatch."
func IsMatch(s string, p string) bool {
	sIdx, pIdx := 0, 0
	star, match := -1, -1

	for sIdx < len(s) {
		if pIdx < len(p) && (p[pIdx] == '?' || p[pIdx] == s[sIdx]) {
			sIdx++
			pIdx++
		} else if pIdx < len(p) && p[pIdx] == '*' {
			star = pIdx
			match = sIdx
			pIdx++
		} else if star != -1 {
			pIdx = star + 1
			match++
			sIdx = match
		} else {
			return false
		}
	}
	for pIdx < len(p) && p[pIdx] == '*' {
		pIdx++
	}
	return pIdx == len(p)
}

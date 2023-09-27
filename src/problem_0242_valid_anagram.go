// https://leetcode.cn/problems/valid-anagram/
// 有效的字母异位词
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, c := range s {
		cnt[c]++
	}
	for _, c := range t {
		cnt[c]--
		if cnt[c] < 0 {
			return false
		}
	}
	return true
}

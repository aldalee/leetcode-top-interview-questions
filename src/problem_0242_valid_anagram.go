// https://leetcode.cn/problems/valid-anagram/
// 有效的字母异位词
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range t {
		if cnt[c-'a']--; cnt[c-'a'] < 0 {
			return false
		}
	}
	return true
}

// https://leetcode.cn/problems/first-unique-character-in-a-string/
// 字符串中的第一个唯一字符
package main

func firstUniqChar(s string) int {
	cnts := make([]int, 26)
	for _, ch := range s {
		cnts[ch-'a']++
	}
	for i, ch := range s {
		if cnts[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

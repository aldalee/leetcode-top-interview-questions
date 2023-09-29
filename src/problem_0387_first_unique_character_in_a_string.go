// https://leetcode.cn/problems/first-unique-character-in-a-string/
// 字符串中的第一个唯一字符
package main

import "math"

func firstUniqChar(s string) int {
	pos := map[rune]int{}
	for i, c := range s {
		if _, ok := pos[c]; ok {
			pos[c] = -1
		} else {
			pos[c] = i
		}
	}
	ans := math.MaxInt32
	for _, idx := range pos {
		if idx != -1 && idx < ans {
			ans = idx
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

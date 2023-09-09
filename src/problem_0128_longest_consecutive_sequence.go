// https://leetcode.cn/problems/longest-consecutive-sequence/
// 最长连续序列
package main

func longestConsecutive(nums []int) (maxLen int) {
	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num]++
			preLen, posLen := m[num-1], m[num+1]
			curLen := preLen + posLen + 1
			m[num-preLen] = curLen
			m[num+posLen] = curLen
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return
}
